package main

import (
	"fmt"
	"testing"
)

func TestClimbing(t *testing.T) {
	testcases := []struct {
		arg  int
		want int
	}{
		{
			arg:  2,
			want: 2,
		},
		{
			arg:  3,
			want: 3,
		},
		{
			arg:  5,
			want: 8,
		},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := climbStairs(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
