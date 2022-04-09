package main

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	testcases := []struct {
		arg  int
		want int
	}{
		{
			arg:  4,
			want: 2,
		},
		{
			arg:  8,
			want: 2,
		},
		{
			arg:  9,
			want: 3,
		},
		{
			arg:  15,
			want: 3,
		},
		{
			arg:  2147483647,
			want: 46340,
		},
		{
			arg:  0,
			want: 0,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := mySqrt(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
