package main

import (
	"fmt"
	"testing"
)

func TestRemvoeElements(t *testing.T) {
	testcases := []struct {
		head *ListNode
		val  int
		want *ListNode
	}{
		{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 6,
									},
								},
							},
						},
					},
				},
			},
			6,
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
		},
		{
			nil,
			1,
			nil,
		},
		{
			&ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 7,
						},
					},
				},
			},
			7,
			nil,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := removeElements(tc.head, tc.val)
			if !isEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
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
