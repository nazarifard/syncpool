package syncpool

import (
	"bytes"
	"encoding/json"
	"testing"

	gobufferpool "github.com/libp2p/go-buffer-pool"
	"github.com/valyala/bytebufferpool"
)

func BenchmarkByteSlicePool(b *testing.B) {
	var pool SlicePool[byte]
	for i := 0; i < b.N; i++ {
		buf := pool.Get(32)
		copy(buf, "abcdefghijklmnopqrstvuwxyz123456")
		pool.Put(buf)
	}
}

func BenchmarkGoBufferPool(b *testing.B) {
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf = gobufferpool.Get(32)
		copy(buf, "abcdefghijklmnopqrstvuwxyz123456")
		gobufferpool.Put(buf)
	}
}

func BenchmarkByteBufferPool(b *testing.B) {
	var bb *bytebufferpool.ByteBuffer
	for i := 0; i < b.N; i++ {
		bb = bytebufferpool.Get()
		_, _ = bb.Write([]byte("abcdefghijklmnopqrstvuwxyz123456"))
		bytebufferpool.Put(bb)
	}
}

func BenchmarkUberZapBuffer(b *testing.B) {
	pool := NewBufferPool()
	for i := 0; i < b.N; i++ {
		zb := pool.Get(0)
		_, _ = zb.Write([]byte("abcdefghijklmnopqrstvuwxyz123456"))
		_ = zb.Bytes()
		zb.Free()
	}
}

func BenchmarkJsonEncoderWithPool(b *testing.B) {
	var wBuff bytes.Buffer
	encoder := json.NewEncoder(&wBuff)
	pool := NewBufferPool()
	for i := 0; i < b.N; i++ {
		wBuff.Reset()
		_ = encoder.Encode("abcdefghijklmnopqrstvuwxyz123456")

		zb := pool.Get(0)
		_, _ = zb.Write(wBuff.Bytes()) //export result
		zb.Free()
	}
}

func BenchmarkJsonMarshalWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal("abcdefghijklmnopqrstvuwxyz123456")
	}
}
