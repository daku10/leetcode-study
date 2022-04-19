package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		argHead  *ListNode
		argLeft  int
		argRight int
		want     *ListNode
	}{
		{
			newListNode([]int{1, 2, 3, 4, 5}),
			2,
			4,
			newListNode([]int{1, 4, 3, 2, 5}),
		},
		{
			newListNode([]int{5}),
			1,
			1,
			newListNode([]int{5}),
		},
		{
			newListNode([]int{1, 2}),
			1,
			2,
			newListNode([]int{2, 1}),
		},
		{
			newListNode([]int{1, 2, 3, 4, 5}),
			4,
			5,
			newListNode([]int{1, 2, 3, 5, 4}),
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argHead, tc.argLeft, tc.argRight), func(t *testing.T) {
			actual := reverseBetween(tc.argHead, tc.argLeft, tc.argRight)
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
	head := &ListNode{
		Val: nums[0],
	}
	current := head
	for i := 1; i < len(nums); i++ {
		tmp := &ListNode{
			Val: nums[i],
		}
		current.Next = tmp
		current = current.Next
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
