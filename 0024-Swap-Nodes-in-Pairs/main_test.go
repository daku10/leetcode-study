package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {

	type TestCase struct {
		arg *ListNode
		expect *ListNode
	}

	testCases := []TestCase{
		{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}}},
		{&ListNode{1, nil}, &ListNode{1, nil}},
		{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{6, &ListNode{5, &ListNode{7, nil}}}}}}}},
		{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{6, &ListNode{5, nil}}}}}}},
		{nil, nil},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := swapPairs(testCase.arg)
			if !isSame(actual, testCase.expect) {
				t.Errorf("actual: %v, expect: %v", actual, testCase.expect)
			}
		})
	}
}

func isSame(l1 *ListNode, l2 *ListNode) bool {

	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	for l1 != nil {
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}

	return reflect.DeepEqual(arr1, arr2)
}