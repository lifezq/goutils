// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type testDef_v1 struct {
	Name string
	Age  uint8
}

type testDef_v2 struct {
	Email   string
	Address string
	Phone   uint8
}

var buf bytes.Buffer

func TestGobEncode(t *testing.T) {
	enc := gob.NewEncoder(&buf)

	for i := 0; i < 100; i++ {

		if err := enc.Encode(testDef_v1{
			Name: fmt.Sprintf("%d.%s", i, "name_string"),
			Age:  uint8(i),
		}); err != nil {
			t.Errorf("Encode : %s\n", err.Error())
		}
	}

	for i := 0; i < 100; i++ {

		if err := enc.Encode(testDef_v2{
			Email:   fmt.Sprintf("%d.%s", i, "mail_string@test.com"),
			Address: fmt.Sprintf("%d.%s", i, "address_string"),
			Phone:   uint8(i),
		}); err != nil {
			t.Errorf("Encode : %s\n", err.Error())
		}
	}

}

func TestGobDecode(t *testing.T) {

	var (
		v1 testDef_v1
		v2 testDef_v2
	)

	dec := gob.NewDecoder(&buf)

	for i := 0; i < 100; i++ {

		if err := dec.Decode(&v1); err != nil {
			t.Errorf("Decode : %s\n", err.Error())
			continue
		}

		t.Logf("V1.Decode : %v\n", v1)
	}

	for i := 0; i < 100; i++ {

		if err := dec.Decode(&v2); err != nil {
			t.Errorf("Decode : %s\n", err.Error())
			continue
		}

		t.Logf("V2.Decode : %v\n", v2)
	}
}
