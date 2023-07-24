package main

import (
	"fmt"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	testcases := []struct {
		matrix [][]byte
		want   int
	}{
		// {
		// 	[][]byte{
		// 		{'1', '0', '1', '0', '0'},
		// 		{'1', '0', '1', '1', '1'},
		// 		{'1', '1', '1', '1', '1'},
		// 		{'1', '0', '0', '1', '0'},
		// 	},
		// 	4,
		// },
		// {
		// 	[][]byte{
		// 		{'0', '1'},
		// 		{'1', '0'},
		// 	},
		// 	1,
		// },
		// {
		// 	[][]byte{
		// 		{'0'},
		// 	},
		// 	0,
		// },
		// {
		// 	[][]byte{
		// 		{'1', '0', '1', '0', '0'},
		// 		{'1', '0', '1', '1', '1'},
		// 		{'1', '1', '1', '1', '1'},
		// 		{'1', '0', '0', '1', '0'},
		// 	},
		// 	4,
		// },
		{
			[][]byte{
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
			},
			9,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maximalSquare(tc.matrix)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
