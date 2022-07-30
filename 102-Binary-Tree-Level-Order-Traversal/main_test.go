package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {

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
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			&TreeNode{
				Val: 1,
			},
			[][]int{{1}},
		},
		{
			nil, [][]int{},
		},
		{
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 3,
					},
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
			[][]int{{3}, {9, 20}, {3, 15, 7}},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := levelOrder(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v arg: %v", got, tc.want, tc.arg)
			}
		})
	}
}
