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
			newListNode([]int{1, 1, 2}),
			newListNode([]int{1, 2}),
		},
		{
			newListNode([]int{1, 1, 2, 3, 3}),
			newListNode([]int{1, 2, 3}),
		},
		{
			newListNode([]int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4}),
			newListNode([]int{1, 2, 3, 4}),
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
	node := &ListNode{
		Val: nums[0],
	}
	head := node
	for i := 1; i < len(nums); i++ {
		tmp := &ListNode{
			Val: nums[i],
		}
		node.Next = tmp
		node = tmp
	}
	return head
}

func toIntSlice(node *ListNode) []int {
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
