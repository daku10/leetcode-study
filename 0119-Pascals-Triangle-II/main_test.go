package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	testcases := []struct {
		arg  int
		want []int
	}{
		{
			3,
			[]int{1, 3, 3, 1},
		},
		{
			0,
			[]int{1},
		},
		{
			1,
			[]int{1, 1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getRow(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
