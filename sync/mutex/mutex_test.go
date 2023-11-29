// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex_Lock(t *testing.T) {
	m := NewMutex()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			lockKey := "foo"
			t.Log(fmt.Sprintf("%d.lock...", i))
			m.Lock(lockKey)
			t.Log(fmt.Sprintf("%d.unlock successful", i))
			time.Sleep(time.Second * 2)
			m.Unlock(lockKey)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
