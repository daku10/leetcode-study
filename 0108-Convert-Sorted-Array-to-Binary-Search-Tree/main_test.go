package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	testcases := []struct {
		arg  []int
		want *TreeNode
	}{
		{
			[]int{-10, -3, 0, 5, 9},
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			[]int{1, 3},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortedArrayToBST(tc.arg)
			if !cmpTree(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}

func cmpTree(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if x.Val != y.Val {
		return false
	}
	return cmpTree(x.Left, y.Left) && cmpTree(x.Right, y.Right)
}
