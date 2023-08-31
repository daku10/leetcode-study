package main

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	type arg struct {
		node *ListNode
		head *ListNode
		want *ListNode
	}

	testcases := []arg{
		func() arg {
			node := &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			}
			head := &ListNode{
				Val:  4,
				Next: node,
			}
			return arg{
				node: node,
				head: head,
				want: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			}
		}(),
		func() arg {
			node := &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			}
			head := &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: node,
				},
			}
			return arg{
				node: node,
				head: head,
				want: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			}
		}(),
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			deleteNode(tc.node)
			if !isEqual(tc.head, tc.want) {
				t.Errorf("head: %v want: %v", tc.head, tc.want)
			}
		})
	}
}

func isEqual(x *ListNode, y *ListNode) bool {
	for {
		if x == nil || y == nil {
			return x == nil && y == nil
		}
		if x.Val != y.Val {
			return false
		}
		x = x.Next
		y = y.Next
	}
}
