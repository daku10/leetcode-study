package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {

	type TestCase struct {
		argHead *ListNode
		argN int
		expect *ListNode
	}

	testCases := []TestCase{
		{
			argHead: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			argN: 2,
			expect: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},		
		},
		{
			argHead: &ListNode{1, nil},
			argN: 1,
			expect: nil,
		},
		{
			argHead: &ListNode{1, &ListNode{2, nil}},
			argN: 1,
			expect: &ListNode{1, nil},
		},
		{
			argHead: &ListNode{1, &ListNode{2, nil}},
			argN: 2,
			expect: &ListNode{2, nil},
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := removeNthFromEnd(testCase.argHead, testCase.argN)
			if !isSame(actual, testCase.expect) {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
				fmt.Println("actual")
				show(actual)
				fmt.Println("expect")
				show(testCase.expect)
			}
		})
	}
}

func show(l *ListNode) {
	for l != nil {
		fmt.Printf("%v ", l.Val)
		l = l.Next
	}
	fmt.Println("")
}

func isSame(l1 *ListNode, l2 *ListNode) bool {
	l1Array := make([]int, 0)
	l2Array := make([]int, 0)

	l1Current := l1
	for l1Current != nil {
		l1Array = append(l1Array, l1Current.Val)
		l1Current = l1Current.Next
	}

	l2Current := l2
	for l2Current != nil {
		l2Array = append(l2Array, l2Current.Val)
		l2Current = l2Current.Next
	}

	return reflect.DeepEqual(l1Array, l2Array)
}
