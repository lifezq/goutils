// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package goutils

import "testing"

func TestDiskRelativePath(t *testing.T) {

	var tdata = []struct {
		p1   string
		p2   string
		want string
	}{
		{
			p1:   "/opt/a/b/c/d/e.go",
			p2:   "/opt/a/f/j/h/i.go",
			want: "../../../b/c/d/e.go",
		},
		{
			p1:   "/opt/a/b/c/d/e.go",
			p2:   "/opt/a/f/j/h/i/j/k/m.go",
			want: "../../../../../../b/c/d/e.go",
		},
		{
			p1:   "/opt/a/b/c/d/e.go",
			p2:   "/opt/a/c/b.go",
			want: "../b/c/d/e.go",
		},
		{
			p1:   "/opt/a/b/c/d/e.go",
			p2:   "/opt/b.go",
			want: "a/b/c/d/e.go",
		},
		{
			p1:   "/opt/a/b/c/d/e.go",
			p2:   "/o/p/b.go",
			want: "../../opt/a/b/c/d/e.go",
		},
	}

	for _, td := range tdata {

		if want := GetRelativePath(td.p1, td.p2); want != td.want {
			t.Errorf("want:%s got:%s\n", td.want, want)
		}
	}

}
