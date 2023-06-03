package main

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	testcases := []struct {
		arg  [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numIslands(tc.arg)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
