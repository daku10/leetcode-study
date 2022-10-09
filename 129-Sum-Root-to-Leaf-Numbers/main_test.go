package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want int
	}{
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
			25,
		},
		{
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			1026,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := sumNumbers(tc.arg)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v arg: %v", actual, tc.want, tc.arg)
			}
		})
	}
}
