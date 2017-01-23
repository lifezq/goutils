// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"encoding/json"
	"os"
)

func JsonEncodeToFile(v interface{}, f string) error {

	fp, err := os.OpenFile(f, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	defer fp.Close()
	if err != nil {
		return err
	}

	return json.NewEncoder(fp).Encode(v)
}

func JsonDecodeFile(f string, v interface{}) error {

	fp, err := os.Open(f)
	defer fp.Close()
	if err != nil {
		return err
	}

	return json.NewDecoder(fp).Decode(&v)
}
