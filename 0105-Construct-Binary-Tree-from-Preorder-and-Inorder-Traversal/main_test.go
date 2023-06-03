package main

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testcases := []struct {
		argPreorder []int
		argInorder  []int
		want        *TreeNode
	}{
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
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
		},
		{
			[]int{-1},
			[]int{-1},
			&TreeNode{
				Val: -1,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := buildTree(tc.argPreorder, tc.argInorder)
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
