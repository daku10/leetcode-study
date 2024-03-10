package main

import (
	"fmt"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	testcases := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
			&ListNode{
				1,
				&ListNode{
					3,
					&ListNode{
						5,
						&ListNode{
							2,
							&ListNode{
								4,
								nil,
							},
						},
					},
				},
			},
		},
		{
			&ListNode{
				2,
				&ListNode{
					1,
					&ListNode{
						3,
						&ListNode{
							5,
							&ListNode{
								6,
								&ListNode{
									4,
									&ListNode{
										7,
										nil,
									},
								},
							},
						},
					},
				},
			},
			&ListNode{
				2,
				&ListNode{
					3,
					&ListNode{
						6,
						&ListNode{
							7,
							&ListNode{
								1,
								&ListNode{
									5,
									&ListNode{
										4,
										nil,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			&ListNode{1, nil},
			&ListNode{1, nil},
		},
		{
			nil,
			nil,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := oddEvenList(tc.head)
			if !isEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", show(got), show(tc.want))
			}
		})
	}
}

func isEqual(a *ListNode, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func show(n *ListNode) string {
	var result string
	for n != nil {
		result += fmt.Sprint(n.Val)
		result += ", "
		n = n.Next
	}
	return result
}
