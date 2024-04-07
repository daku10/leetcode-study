package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	testcases := []struct {
		n    int
		want []int
	}{
		{
			2,
			[]int{0, 1, 1},
		},
		{
			5,
			[]int{0, 1, 1, 2, 1, 2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countBits(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
