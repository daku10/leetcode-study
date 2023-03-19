package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want []int
	}{
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			[]int{3, 2, 1},
		},
		{
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			[]int{1, 2, 3},
		},
		{
			nil,
			nil,
		},
		{
			&TreeNode{Val: 1},
			[]int{1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := postorderTraversal(tc.arg)
			if !reflect.DeepEqual(actual, tc.want) {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
