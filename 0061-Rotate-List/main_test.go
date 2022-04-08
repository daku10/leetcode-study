package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testcases := []struct {
		argHead *ListNode
		argK    int
		want    *ListNode
	}{
		{
			argHead: newListNode([]int{1, 2, 3, 4, 5}),
			argK:    2,
			want:    newListNode([]int{4, 5, 1, 2, 3}),
		},
		{
			argHead: newListNode([]int{0, 1, 2}),
			argK:    4,
			want:    newListNode([]int{2, 0, 1}),
		},
		{
			argHead: nil,
			argK:    5,
			want:    nil,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argHead), func(t *testing.T) {
			actual := rotateRight(tc.argHead, tc.argK)
			if !reflect.DeepEqual(toSlice(tc.want), toSlice(actual)) {
				t.Errorf("expected: %v, actual: %v", toSlice(tc.want), toSlice(actual))
			}
		})
	}
}

func toSlice(node *ListNode) []int {
	result := make([]int, 0)
	n := node
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	return result
}

func newListNode(val []int) *ListNode {
	result := &ListNode{
		Val: val[0],
	}
	current := result
	for i := 0; i < len(val)-1; i++ {
		next := &ListNode{Val: val[i+1]}
		current.Next = next
		current = next
	}
	return result
}
