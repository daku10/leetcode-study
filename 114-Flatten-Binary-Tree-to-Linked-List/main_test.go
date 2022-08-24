package main

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want *TreeNode
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
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
			},
		},
		{
			nil,
			nil,
		},
		{
			&TreeNode{
				Val: 0,
			},
			&TreeNode{
				Val: 0,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			flatten(tc.arg)
			if !cmpTree(tc.want, tc.arg) {
				t.Errorf("got: %v want: %v", tc.arg, tc.want)
				printPreorder(tc.arg)
				fmt.Println()
				printPreorder(tc.want)
			}
		})
	}
}

func cmpTree(x, y *TreeNode) bool {
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

func printPreorder(x *TreeNode) {
	if x == nil {
		return
	}
	fmt.Println(x.Val)
	printPreorder(x.Left)
	printPreorder(x.Right)
}
