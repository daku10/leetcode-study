package main

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	testcases := []struct {
		arg  string
		want int
	}{
		{
			arg:  "Hello World",
			want: 5,
		},
		{
			arg:  "   fly me   to   the moon  ",
			want: 4,
		},
		{
			arg:  "luffy is still joyboy",
			want: 6,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := lengthOfLastWord(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
