package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {

	testcases := []struct {
		argHead *ListNode
		argX    int
		want    *ListNode
	}{
		{
			newListNode([]int{1, 4, 3, 2, 5, 2}),
			3,
			newListNode([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			newListNode([]int{2, 1}),
			2,
			newListNode([]int{1, 2}),
		},
		{
			newListNode([]int{4, 3, 2, 5, 2}),
			3,
			newListNode([]int{2, 2, 4, 3, 5}),
		},
		{
			newListNode([]int{1, 1}),
			2,
			newListNode([]int{1, 1}),
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argHead, tc.argX), func(t *testing.T) {
			actual := partition(tc.argHead, tc.argX)
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
		// fmt.Println(node.Val)
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
