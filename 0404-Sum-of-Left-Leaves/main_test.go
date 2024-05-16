package main

import (
	"fmt"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
	testcases := []struct {
		root *TreeNode
		want int
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
			24,
		},
		{
			&TreeNode{
				Val: 1,
			},
			0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sumOfLeftLeaves(tc.root)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
