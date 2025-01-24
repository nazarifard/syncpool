package syncpool

import zapbuffer "go.uber.org/zap/buffer"

type BufferPool zapbuffer.Pool
type Buffer struct{ *zapbuffer.Buffer }

func NewBufferPool() BufferPool {
	return BufferPool(zapbuffer.NewPool())
}

var zeroBuffer [1024]byte

func (bp BufferPool) Get(size int) Buffer {
	zb := zapbuffer.Pool(bp).Get()
	if size == 0 {
		return Buffer{zb}
	}
	if size > len(zeroBuffer) {
		for i := 0; i < size/len(zeroBuffer); i++ {
			zb.Write(zeroBuffer[:])
		}
		zb.Write(zeroBuffer[:size%len(zeroBuffer)])
		return Buffer{zb}
	}
	zb.Write(zeroBuffer[:size])
	return Buffer{zb}
}

func (b *Buffer) Read(bs []byte) (n int, err error) {
	n = copy(bs, b.Bytes())
	return
}

//Buffer
//usage
//p:=NewBufferPool()
//buf:=p.Get()
//dosomthing(buf)
//buf.Free()
