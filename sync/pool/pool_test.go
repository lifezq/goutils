// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package pool

import (
	"io"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := New(newConnector, 5)

	w := sync.WaitGroup{}
	poolSize := 10
	w.Add(poolSize)
	for i := 1; i <= poolSize; i++ {
		go func() {
			defer w.Done()

			conn, err := p.Pull()
			if err != nil {
				t.Errorf("p.Pull %s\n", err.Error())
				return
			}

			t.Logf("connector id:%v\n", conn.(*Conn).Get())

			p.Push(conn)
		}()

		if i > poolSize/2 {
			time.Sleep(time.Millisecond * 10)
		}
	}

	w.Wait()

	p.Close()
}

type Conn struct {
	id int32
}

var ider int32

func newConnector() (io.Closer, error) {
	return &Conn{
		id: atomic.AddInt32(&ider, 1),
	}, nil
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) Get() int32 {
	return c.id
}
