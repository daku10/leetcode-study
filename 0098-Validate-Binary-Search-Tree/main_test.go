package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want bool
	}{
		{
			arg: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: true,
		},
		{
			arg: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: false,
		},
		{
			arg: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			want: true,
		},
		{
			arg: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: false,
		},
		{
			arg: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			want: false,
		},
		{
			arg: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%v", tc.arg), func(t *testing.T) {
			got := isValidBST(tc.arg)
			if tc.want != got {
				t.Errorf("failed arg=%v want=%v got=%v", tc.arg, tc.want, got)
			}
		})
	}
}
