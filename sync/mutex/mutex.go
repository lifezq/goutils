// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package mutex

import "sync"

type Mutex map[string]*sync.Mutex

func NewMutex() *Mutex {
	m := Mutex(make(map[string]*sync.Mutex))
	return &m
}

func (m *Mutex) Lock(key string) {
	if _, ok := (*m)[key]; !ok {
		(*m)[key] = &sync.Mutex{}
	}
	(*m)[key].Lock()
}

func (m *Mutex) Unlock(key string) {
	if _, ok := (*m)[key]; !ok {
		(*m)[key] = &sync.Mutex{}
		return
	}
	(*m)[key].Unlock()
}
