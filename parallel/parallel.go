// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package parallel

func Parallel(fns ...func()) {
	if len(fns) == 0 {
		return
	}
	group := NewRoutineGroup()
	for _, fn := range fns {
		group.RunSafe(fn)
	}
	group.Wait()
}
