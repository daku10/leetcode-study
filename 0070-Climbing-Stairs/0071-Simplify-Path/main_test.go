package main

import (
	"fmt"
	"testing"
)

func TestSimplifiedPath(t *testing.T) {
	testcases := []struct {
		arg  string
		want string
	}{
		{
			arg:  "/home/",
			want: "/home",
		},
		{
			arg:  "/../",
			want: "/",
		},
		{
			arg:  "/home//foo",
			want: "/home/foo",
		},
		{
			arg:  "/home/foo/../../hoge",
			want: "/hoge",
		},
		{
			arg:  "/../../",
			want: "/",
		},
		{
			arg:  "/../../../fuga",
			want: "/fuga",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := simplifyPath(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
