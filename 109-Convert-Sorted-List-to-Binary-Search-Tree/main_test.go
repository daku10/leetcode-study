package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	testcases := []struct {
		arg  *ListNode
		want *TreeNode
	}{
		{
			&ListNode{
				-10,
				&ListNode{
					-3,
					&ListNode{
						0,
						&ListNode{
							5,
							&ListNode{
								9,
								nil,
							},
						},
					},
				},
			},
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
			&ListNode{
				1,
				&ListNode{
					3,
					nil,
				},
			},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			nil,
			nil,
		},
		{
			&ListNode{
				0,
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							&ListNode{
								4,
								&ListNode{5, nil},
							},
						},
					},
				},
			},
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
			got := sortedListToBST(tc.arg)
			if !cmpTree(got, tc.want) {
				printTree(got)
				fmt.Println()
				printTree(tc.want)
				fmt.Println()
				t.Errorf("got: %#v want: %#v", got, tc.want)
			}
		})
	}
}

func printTree(t *TreeNode) {
	if t == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Printf("Val: %d ", t.Val)
	fmt.Printf("Left: ")
	printTree(t.Left)
	fmt.Printf("Right: ")
	printTree(t.Right)
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
