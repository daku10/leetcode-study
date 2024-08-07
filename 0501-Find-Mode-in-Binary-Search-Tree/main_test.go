package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMode(t *testing.T) {
	testcases := []struct {
		root *TreeNode
		want []int
	}{
		{
			&TreeNode{
				1,
				nil,
				&TreeNode{
					2,
					&TreeNode{
						2,
						nil,
						nil,
					},
					nil,
				},
			},
			[]int{2},
		},
		{
			&TreeNode{
				0, nil, nil,
			},
			[]int{0},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findMode(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
