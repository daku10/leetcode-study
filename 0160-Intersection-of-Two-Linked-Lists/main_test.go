package main

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	testcases := []struct {
		argA *ListNode
		argB *ListNode
		want *ListNode
	}{
		func() struct {
			argA *ListNode
			argB *ListNode
			want *ListNode
		} {
			a1 := &ListNode{
				Val: 4,
			}
			a2 := &ListNode{
				Val: 1,
			}
			b1 := &ListNode{
				Val: 5,
			}
			b2 := &ListNode{
				Val: 6,
			}
			b3 := &ListNode{
				Val: 1,
			}
			c1 := &ListNode{
				Val: 8,
			}
			c2 := &ListNode{
				Val: 4,
			}
			c3 := &ListNode{
				Val: 5,
			}
			a1.Next = a2
			a2.Next = c1
			b1.Next = b2
			b2.Next = b3
			b3.Next = c1
			c1.Next = c2
			c2.Next = c3

			return struct {
				argA *ListNode
				argB *ListNode
				want *ListNode
			}{
				a1,
				b1,
				c1,
			}
		}(),
		func() struct {
			argA *ListNode
			argB *ListNode
			want *ListNode
		} {
			b1 := &ListNode{
				Val: 5,
			}
			b2 := &ListNode{
				Val: 6,
			}
			b3 := &ListNode{
				Val: 1,
			}
			c1 := &ListNode{
				Val: 8,
			}
			c2 := &ListNode{
				Val: 4,
			}
			c3 := &ListNode{
				Val: 5,
			}
			b1.Next = b2
			b2.Next = b3
			b3.Next = c1
			c1.Next = c2
			c2.Next = c3

			return struct {
				argA *ListNode
				argB *ListNode
				want *ListNode
			}{
				c1,
				b1,
				c1,
			}
		}(),
		func() struct {
			argA *ListNode
			argB *ListNode
			want *ListNode
		} {
			a1 := &ListNode{
				Val: 4,
			}
			a2 := &ListNode{
				Val: 1,
			}
			b1 := &ListNode{
				Val: 5,
			}
			b2 := &ListNode{
				Val: 6,
			}
			b3 := &ListNode{
				Val: 1,
			}
			a1.Next = a2
			b1.Next = b2
			b2.Next = b3

			return struct {
				argA *ListNode
				argB *ListNode
				want *ListNode
			}{
				a1,
				b1,
				nil,
			}
		}(),
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := getIntersectionNode(tc.argA, tc.argB)
			if actual != tc.want {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
