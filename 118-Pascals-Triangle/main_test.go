package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testcases := []struct {
		arg  int
		want [][]int
	}{
		{
			5,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			1,
			[][]int{
				{1},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := generate(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
