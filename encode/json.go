// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"encoding/json"
	"os"
)

func JSONEncodeToFile(v interface{}, f string) error {
	fp, err := os.OpenFile(f, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644) //nolint:gocritic // ignored
	defer fp.Close()                                                    //nolint:staticcheck // ignored
	if err != nil {
		return err
	}

	return json.NewEncoder(fp).Encode(v)
}

func JSONDecodeFile(f string, v interface{}) error {
	fp, err := os.Open(f)
	defer fp.Close() //nolint:staticcheck // ignored
	if err != nil {
		return err
	}

	return json.NewDecoder(fp).Decode(&v)
}
