package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	testcases := []struct {
		s    string
		want string
	}{
		{
			s:    "hello",
			want: "olleh",
		},
		{
			s:    "Hannah",
			want: "hannaH",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			arg := []byte(tc.s)
			reverseString(arg)
			if string(arg) != tc.want {
				t.Errorf("got: %v, want: %v", string(arg), tc.want)
			}
		})
	}
}
