package main

import (
	"fmt"
	"testing"
)

func TestMinDepth(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
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
			2,
		},
		{
			&TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			5,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minDepth(tc.arg)
			if got != tc.want {
				t.Errorf("failed arg: %v got: %v want: %v", tc.arg, got, tc.want)
			}
		})
	}
}
