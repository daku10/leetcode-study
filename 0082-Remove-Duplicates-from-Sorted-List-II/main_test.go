package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	testcases := []struct {
		arg  *ListNode
		want *ListNode
	}{
		{
			newListNode([]int{1, 2, 3, 3, 4, 4, 5}),
			newListNode([]int{1, 2, 5}),
		},
		{
			newListNode([]int{1, 1, 1, 2, 3}),
			newListNode([]int{2, 3}),
		},
		{
			newListNode([]int{1, 1}),
			nil,
		},
		{
			nil,
			nil,
		},
		{
			newListNode([]int{1, 1, 1, 1}),
			nil,
		},
		{
			newListNode([]int{1, 1, 2, 2}),
			nil,
		},
		{
			newListNode([]int{2, 3, 4, 4}),
			newListNode([]int{2, 3}),
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := deleteDuplicates(tc.arg)
			if !reflect.DeepEqual(toIntSlice(tc.want), toIntSlice(actual)) {
				t.Errorf("expected: %v, actual: %v", toIntSlice(tc.want), toIntSlice(actual))
			}
		})
	}
}

func newListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	result := head
	for i := 1; i < len(nums); i++ {
		tmp := &ListNode{Val: nums[i], Next: nil}
		head.Next = tmp
		head = tmp
	}
	return result
}

func toIntSlice(node *ListNode) []int {
	result := make([]int, 0)
	current := node
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
