package main

import (
	"fmt"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	testcases := []struct {
		arg  *ListNode
		want *ListNode
	}{
		{
			&ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := insertionSortList(tc.arg)
			if !isEqual(actual, tc.want) {
				t.Errorf("actual: %v want: %v", convertList(actual), convertList(tc.want))
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

func convertList(n *ListNode) []int {
	var result []int
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	return result
}
