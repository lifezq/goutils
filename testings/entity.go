// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package testings

import "testing"

type Entity struct {
	Args any
	Want any
	Name string
}

type Entities []*Entity

func NewEntities(params ...*Entity) *Entities {
	data := make(Entities, 0)
	for _, param := range params {
		if param != nil {
			data = append(data, param)
		}
	}
	return &data
}

func (e *Entity) Equal(result any) bool {
	return result == e.Want
}

//nolint:thelper // ignore
func (e *Entities) Run(t *testing.T) {
	for _, entity := range *e {
		if entity == nil {
			t.Fatal("invalid memory address or nil pointer dereference")
		}
		t.Run(entity.Name, func(t *testing.T) {
			if !entity.Equal(entity.Args) {
				t.Errorf("%s: %v!= %v", entity.Name, entity.Args, entity.Want)
			}
		})
	}
}
