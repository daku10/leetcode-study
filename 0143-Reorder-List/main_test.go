package main

import (
	"fmt"
	"testing"
)

func TestReorderList(t *testing.T) {
	testcases := []struct {
		arg  *ListNode
		want *ListNode
	}{
		{
			(func() *ListNode {
				n0 := &ListNode{
					Val: 1,
				}
				n1 := &ListNode{
					Val: 2,
				}
				n2 := &ListNode{
					Val: 3,
				}
				n3 := &ListNode{
					Val: 4,
				}
				n0.Next = n1
				n1.Next = n2
				n2.Next = n3
				return n0
			})(),
			(func() *ListNode {
				n0 := &ListNode{
					Val: 1,
				}
				n1 := &ListNode{
					Val: 2,
				}
				n2 := &ListNode{
					Val: 3,
				}
				n3 := &ListNode{
					Val: 4,
				}
				n0.Next = n3
				n3.Next = n1
				n1.Next = n2
				return n0
			})(),
		},
		{
			(func() *ListNode {
				n0 := &ListNode{
					Val: 1,
				}
				n1 := &ListNode{
					Val: 2,
				}
				n2 := &ListNode{
					Val: 3,
				}
				n3 := &ListNode{
					Val: 4,
				}
				n4 := &ListNode{
					Val: 5,
				}
				n0.Next = n1
				n1.Next = n2
				n2.Next = n3
				n3.Next = n4
				return n0
			})(),
			(func() *ListNode {
				n0 := &ListNode{
					Val: 1,
				}
				n1 := &ListNode{
					Val: 2,
				}
				n2 := &ListNode{
					Val: 3,
				}
				n3 := &ListNode{
					Val: 4,
				}
				n4 := &ListNode{
					Val: 5,
				}
				n0.Next = n4
				n4.Next = n1
				n1.Next = n3
				n3.Next = n2
				return n0
			})(),
		},
		{
			&ListNode{
				Val: 1,
			},
			&ListNode{
				Val: 1,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			reorderList(tc.arg)
			if !isEqual(tc.arg, tc.want) {
				t.Errorf("actual: %v want: %v", tc.arg, tc.want)
			}
		})
	}
}

func isEqual(a *ListNode, b *ListNode) bool {
	for {
		if a == nil || b == nil {
			return a == nil && b == nil
		}
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
}
