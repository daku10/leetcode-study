// func numTrees(n int) int {

// }

package main

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	testcases := []struct {
		arg  int
		want int
	}{
		{
			3, 5,
		},
		{
			1, 1,
		},
		{
			2, 2,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := numTrees(tc.arg)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
