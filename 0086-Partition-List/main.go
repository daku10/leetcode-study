package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	beforeHead := &ListNode{
		Next: nil,
	}
	var before *ListNode = beforeHead
	afterHead := &ListNode{
		Next: nil,
	}
	var after *ListNode = afterHead
	current := head
	for current != nil {
		if current.Val < x {
			before.Next = current
			before = current
		} else {
			after.Next = current
			after = after.Next
		}

		current = current.Next
	}

	before.Next = afterHead.Next
	after.Next = nil

	return beforeHead.Next
}
