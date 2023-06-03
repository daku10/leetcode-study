package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTraversal(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want [][]int
	}{
		{
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			[][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			&TreeNode{Val: 1},
			[][]int{{1}},
		},
		{
			nil,
			[][]int{},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := levelOrderBottom(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
