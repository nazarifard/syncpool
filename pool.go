package syncpool

import "sync"

type Pool[V any] sync.Pool

func NewPool[V any]() *Pool[V] {
	p := &sync.Pool{
		New: func() any {
			return new(V)
		},
	}
	return (*Pool[V])(p)
}

func (p *Pool[V]) Get() *V {
	v := (*sync.Pool)(p).Get().(*V)
	var zero V
	*v = zero //reset
	return v
}

func (p *Pool[V]) Put(v *V) {
	(*sync.Pool)(p).Put(v)
}
