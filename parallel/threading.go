package parallel

import (
	"log"
	"runtime"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var (
	// globalPool is the global ants pool for managing goroutines
	globalPool *ants.Pool
)

const (
	// defaultPoolSize is the default size of the ants pool
	defaultPoolSize = 1000
)

func init() {
	var err error
	// Create a pool with default size and a panic handler
	globalPool, err = ants.NewPool(defaultPoolSize, ants.WithPanicHandler(func(p interface{}) {
		buf := make([]byte, 10240)
		runtime.Stack(buf, false)
		log.Printf("ants recover, info=%v, err=%v", string(buf), p)
	}))
	if err != nil {
		panic(err)
	}
}

// GetPool returns the global ants pool
func GetPool() *ants.Pool {
	return globalPool
}

type RoutineGroup struct {
	waitGroup sync.WaitGroup
}

func NewRoutineGroup() *RoutineGroup {
	return new(RoutineGroup)
}

func (g *RoutineGroup) Run(fn func()) {
	g.waitGroup.Add(1)

	err := globalPool.Submit(func() {
		defer g.waitGroup.Done()
		fn()
	})

	if err != nil {
		// If submit fails (e.g. pool closed), fallback to standard go routine
		// or just log error. Ideally fallback to ensure execution.
		// Since we use blocking pool, error usually means pool is closed.
		log.Printf("submit to ants pool failed: %v, fallback to go", err)
		go func() {
			defer g.waitGroup.Done()
			fn()
		}()
	}
}

func (g *RoutineGroup) RunSafe(fn func()) {
	g.waitGroup.Add(1)

	GoSafe(func() {
		defer g.waitGroup.Done()
		fn()
	})
}

func (g *RoutineGroup) Parallel(fns ...func()) {
	if len(fns) == 0 {
		return
	}
	for _, fn := range fns {
		g.RunSafe(fn)
	}
}

func GoSafe(fn func()) {
	// Use ants pool to submit the task
	err := globalPool.Submit(func() {
		// We still keep Recover here if we want to be double safe,
		// but ants WithPanicHandler already handles it.
		// However, RunSafe also does `defer Recover()`.
		// If we use ants panic handler, we don't need explicit Recover inside the task
		// UNLESS we want to keep the exact same behavior as before.
		// The previous RunSafe called `Recover`.
		// Let's rely on ants PanicHandler for the pool.
		// BUT wait, if we fallback to `go func` in `Run`, we don't have panic handler there unless we wrap it.
		// `GoSafe` previously called `go RunSafe(fn)`.

		// Let's use RunSafe logic inside the submitted task to ensure consistency
		// and because our ants panic handler just logs, while Recover might do more in future.
		// Also `Recover` in this file logs.

		RunSafe(fn)
	})

	if err != nil {
		log.Printf("submit to ants pool failed: %v, fallback to go safe", err)
		go RunSafe(fn)
	}
}

func RunSafe(fn func()) {
	defer Recover()
	fn()
}

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}
	if p := recover(); p != nil {
		buf := make([]byte, 10240)
		runtime.Stack(buf, false)
		log.Printf("recover, info=%v, err=%v", string(buf), p)
	}
}

// Wait waits all running functions to be done.
func (g *RoutineGroup) Wait() {
	g.waitGroup.Wait()
}
