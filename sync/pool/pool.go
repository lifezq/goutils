package pool

import (
	"errors"
	"io"
	"sync"
)

type Pool struct {
	mux     sync.Mutex
	res     chan io.Closer
	factory func() io.Closer
	closed  bool
}

func New(f func() io.Closer, size uint16) *Pool {

	if size < 1 {
		return nil
	}

	return &Pool{
		factory: f,
		res:     make(chan io.Closer, size),
	}
}

func (p *Pool) Pull() (io.Closer, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	select {
	case r, ok := <-p.res:
		if !ok {
			return nil, errors.New("error")
		}

		return r, nil
	default:
		return p.factory(), nil
	}
}

func (p *Pool) Push(r io.Closer) {
	p.mux.Lock()
	defer p.mux.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
	default:
		r.Close()
	}
}

func (p *Pool) Close() {
	p.mux.Lock()
	defer p.mux.Unlock()

	if p.closed {
		return
	}

	close(p.res)

	for r := range p.res {
		r.Close()
	}

	p.closed = true
}
