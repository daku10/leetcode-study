package main

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	testcases := []struct {
		argNode   *TreeNode
		argTarget int
		want      bool
	}{
		{
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			22,
			true,
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			5,
			false,
		},
		{
			nil,
			0,
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hasPathSum(tc.argNode, tc.argTarget)
			if got != tc.want {
				t.Errorf("got: %v want: %v node: %v target: %v", got, tc.want, tc.argNode, tc.argTarget)
			}
		})
	}
}
