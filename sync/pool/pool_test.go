// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package pool

import (
	"io"
	"sync"
	"sync/atomic"
	"testing"
)

func TestPool(t *testing.T) {
	p := New(newConnector, 5)

	w := sync.WaitGroup{}
	w.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			defer w.Done()

			conn, err := p.Pull()
			if err != nil {
				t.Errorf("p.Pull %s\n", err.Error())
				return
			}

			t.Logf("connector id:%v\n", conn.(*Conn).ID)

			p.Push(conn)
		}()
	}

	w.Wait()

	p.Close()
}

type Conn struct {
	ID int32
}

var ider int32

func newConnector() (io.Closer, error) {
	return &Conn{
		ID: atomic.AddInt32(&ider, 1),
	}, nil
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) Get() int32 {
	return c.ID
}
