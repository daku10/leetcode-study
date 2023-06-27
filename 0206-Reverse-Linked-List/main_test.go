package main

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	testcases := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			&ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
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
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reverseList(tc.head)
			if !isEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}

func isEqual(a *ListNode, b *ListNode) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if a.Val != b.Val {
		return false
	}
	return isEqual(a.Next, b.Next)
}
