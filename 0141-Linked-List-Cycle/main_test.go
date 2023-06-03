package main

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	testcases := []struct {
		arg  *ListNode
		want bool
	}{
		{
			func() *ListNode {
				n0 := &ListNode{
					Val: 3,
				}
				n1 := &ListNode{
					Val: 2,
				}
				n2 := &ListNode{
					Val: 0,
				}
				n3 := &ListNode{
					Val: -4,
				}
				n0.Next = n1
				n1.Next = n2
				n2.Next = n3
				n3.Next = n1
				return n0
			}(),
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := hasCycle(tc.arg)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
