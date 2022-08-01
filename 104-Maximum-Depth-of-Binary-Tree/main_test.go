package main

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want int
	}{
		{
			&TreeNode{
				3,
				&TreeNode{
					Val: 9,
				},
				&TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			3,
		},
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			2,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxDepth(tc.arg)
			if tc.want != got {
				t.Errorf("got: %v want: %v arg: %v", got, tc.want, tc.arg)
			}
		})
	}
}
