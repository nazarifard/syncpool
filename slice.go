package syncpool

import (
	"math"
	"math/bits"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

// base refrence: https://github.com/panjf2000/gnet/blob/v1.6.7/pkg/pool/byteslice/byteslice.go
type SlicePool[T any] struct {
	pools [32]sync.Pool
}

func (p *SlicePool[T]) Get(size int) (buf []T) {
	if size <= 0 {
		return nil
	}
	if size > math.MaxInt32 {
		return make([]T, size)
	}
	idx := index(uint32(size))
	ptr, _ := p.pools[idx].Get().(unsafe.Pointer)
	if ptr == nil {
		return make([]T, 1<<idx)[:size]
	}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	sh.Data = uintptr(ptr)
	sh.Len = size
	sh.Cap = 1 << idx
	runtime.KeepAlive(ptr)
	return
}

func (p *SlicePool[T]) Put(buf []T) {
	size := cap(buf)
	if size == 0 || size > math.MaxInt32 {
		return
	}
	idx := index(uint32(size))
	if size != 1<<idx { // this byte slice is not from Pool.Get(), put it into the previous interval of idx
		idx--
	}
	// array pointer
	p.pools[idx].Put(unsafe.Pointer(&buf[:1][0]))
}

func index(n uint32) uint32 {
	return uint32(bits.Len32(n - 1))
}