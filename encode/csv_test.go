// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package encode

import (
	"bytes"
	"strings"
	"testing"
)

func TestHasBOMAndRemoveBOM(t *testing.T) {
	data := []byte{0xEF, 0xBB, 0xBF, 'a', 'b', 'c'}
	if !hasBOM(data) {
		t.Fatalf("expected BOM")
	}
	trim := removeBOM(data)
	if string(trim) != "abc" {
		t.Fatalf("unexpected removeBOM result: %q", string(trim))
	}
	no := []byte("xyz")
	if hasBOM(no) {
		t.Fatalf("unexpected BOM")
	}
}

func TestParseCSVFromBytes(t *testing.T) {
	src := "a,b,c\n\n1,2,3\n\"x,y\",z\n"
	records, err := parseCSVFromBytes([]byte(src))
	if err != nil {
		t.Fatalf("parseCSVFromBytes error: %v", err)
	}
	if len(records) != 5 {
		t.Fatalf("expected 5 records (including trailing empty), got %d", len(records))
	}
	if strings.Join(records[0], ",") != "a,b,c" {
		t.Fatalf("unexpected header: %v", records[0])
	}
	if len(records[1]) != 0 {
		t.Fatalf("empty line should produce empty record")
	}
}

func TestHasConsistentFields(t *testing.T) {
	consistent := [][]string{
		{"a", "b", "c"},
		{"1", "2", "3"},
		{"x", "y", "z"},
	}
	if !hasConsistentFields(consistent) {
		t.Fatalf("expected consistent")
	}
	inconsistent := [][]string{
		{"a", "b"},
		{"1", "2", "3"},
	}
	if hasConsistentFields(inconsistent) {
		t.Fatalf("expected not consistent")
	}
}

func TestHasReasonableContent(t *testing.T) {
	if hasReasonableContent([][]string{}) {
		t.Fatalf("empty should be false")
	}
	if !hasReasonableContent([][]string{{""}, {"a"}}) {
		t.Fatalf("expected true with non-empty field")
	}
}

func TestForceDecodeCSV_UTF8(t *testing.T) {
	data := "col1,col2\n你好,world\n"
	buf := bytes.NewBufferString(data)
	records, enc, err := ForceDecodeCSV(buf)
	if err != nil {
		t.Fatalf("ForceDecodeCSV error: %v", err)
	}
	if enc != "utf-8" {
		t.Fatalf("expected utf-8, got %s", enc)
	}
	if len(records) < 2 || records[1][0] != "你好" || records[1][1] != "world" {
		t.Fatalf("unexpected records: %v", records)
	}
}
