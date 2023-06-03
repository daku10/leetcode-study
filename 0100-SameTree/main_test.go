package main

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	testcases := []struct {
		arg1 *TreeNode
		arg2 *TreeNode
		want bool
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			true,
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			false,
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isSameTree(tc.arg1, tc.arg2)
			if tc.want != got {
				t.Errorf("got=%v want=%v arg1=%v arg2=%v", got, tc.want, tc.arg1, tc.arg2)
			}
		})
	}
}
