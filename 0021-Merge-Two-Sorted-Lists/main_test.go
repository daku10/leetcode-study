package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	type TestCase struct {
		arg1 *ListNode
		arg2 *ListNode
		expect *ListNode
	}

	testCases := []TestCase{
		{
			arg1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			arg2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			expect: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			arg1: nil,
			arg2: nil,
			expect: nil,
		},
		{
			arg1: nil,
			arg2: &ListNode{0, nil},
			expect: &ListNode{0, nil},
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := mergeTwoLists(testCase.arg1, testCase.arg2)
			if !isSame(actual, testCase.expect) {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
				fmt.Println("actual")
				showList(actual)
				fmt.Println("expect")
				showList(testCase.expect)
			}
		})
	}
}

func isSame(left *ListNode, right *ListNode) bool {
	initLeft := left
	initRight := right
	for {
		if initLeft == nil && initRight == nil {
			return true
		} else if initLeft == nil {
			return false
		} else if initRight == nil {
			return false
		} else if initLeft.Val != initRight.Val {
			return false
		} else {
			initLeft = initLeft.Next
			initRight = initRight.Next
		}
	}
}

func showList(list *ListNode) {
	current := list
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}