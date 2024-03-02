package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	testcases := []struct {
		arg  [][]int
		want [][]int
	}{
		{
			[][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			[][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			[][]int{
				{1, 1},
				{1, 0},
			},
			[][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gameOfLife(tc.arg)
			if !reflect.DeepEqual(tc.arg, tc.want) {
				t.Errorf("got: %v, want: %v", tc.arg, tc.want)
			}
		})
	}
}
