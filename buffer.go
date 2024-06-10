package syncpool

import zapbuffer "go.uber.org/zap/buffer"

type BufferPool zapbuffer.Pool
type Buffer = *zapbuffer.Buffer

func NewBufferPool() BufferPool {
	return BufferPool(zapbuffer.NewPool())
}

var zeroBuffer [1024]byte

func (bp BufferPool) Get(size int) Buffer {
	buffer := zapbuffer.Pool(bp).Get()
	if size == 0 {
		return buffer
	}
	if size > len(zeroBuffer) {
		for i := 0; i < size/len(zeroBuffer); i++ {
			buffer.Write(zeroBuffer[:])
		}
		buffer.Write(zeroBuffer[:size%len(zeroBuffer)])
		return buffer
	}
	buffer.Write(zeroBuffer[:size])
	return buffer
}

//usage
//p:=NewBufferPool()
//buf:=p.Get()
//dosomthing(buf)
//buf.Free()
