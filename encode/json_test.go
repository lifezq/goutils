// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"bytes"
	"encoding/json"
	"testing"
)

type Card struct {
	Name  string
	Alias string
}

var (
	card = Card{
		Name:  "card_name",
		Alias: "alias_value",
	}
)

func TestCompact(t *testing.T) {

	var (
		buf = new(bytes.Buffer)
		src = []byte("{\"name\":    \" name_string  \"  , \"value\" : \" value_string \"  }")
	)

	if err := json.Compact(buf, src); err != nil {
		t.Errorf("Compact Error:%s\n", err.Error())
	}

	t.Logf("buf len:%d:%v\n", len(buf.Bytes()), buf.String())
}

func TestHTMLEscape(t *testing.T) {

	var (
		buf bytes.Buffer
		src = []byte(`
		<html>
		<head><title>test page</title></head>
		<body>
		<p>Hello world!...</p>
		</body>
		</html>
		`)
	)

	json.HTMLEscape(&buf, src)

	t.Logf("buf len:%d:%s\n", len(buf.Bytes()), buf.String())
}

func TestIndent(t *testing.T) {

	src, err := json.Marshal(card)
	if err != nil {
		t.Errorf("Marshal Error:%s\n", err.Error())
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, src, "=", "---"); err != nil {
		t.Errorf("Indent Error:%s\n", err.Error())
	}

	t.Logf("buf len:%d:%s\n", len(buf.Bytes()), buf.String())
}

func TestMarshalIndent(t *testing.T) {

	b, err := json.Marshal(card)
	if err != nil {
		t.Errorf("Marshal Error:%s\n", err.Error())
	}

	t.Logf("len:%d b:%s\n", len(b), string(b))

	b, err = json.MarshalIndent(card, "", "    ")
	if err != nil {
		t.Errorf("MarshalIndent Error:%s\n", err.Error())
	}

	t.Logf("len:%d b:%s\n", len(b), string(b))
}

func TestJsonEncode(t *testing.T) {

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	if err := enc.Encode(card); err != nil {
		t.Errorf("Encode Error:%s\n", err.Error())
	}

	t.Logf("buf:%s\n", buf.String())
}

func TestJsonDecode(t *testing.T) {

	b, err := json.Marshal(card)
	if err != nil {
		t.Errorf("Marshal Error:%s\n", err.Error())
	}

	dec := json.NewDecoder(bytes.NewReader(b))

	var c Card
	if err := dec.Decode(&c); err != nil {
		t.Errorf("Decode Error:%s\n", err.Error())
	}

	t.Logf("Card:%v\n", c)
}
