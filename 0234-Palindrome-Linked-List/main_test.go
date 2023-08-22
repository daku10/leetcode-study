package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		head *ListNode
		want bool
	}{
		{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			true,
		},
		{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPalindrome(tc.head)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
