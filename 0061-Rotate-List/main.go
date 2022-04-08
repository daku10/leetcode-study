package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	current := head
	last := current
	for current != nil {
		if current.Next == nil {
			last = current
		}
		current = current.Next
		length++
	}
	rotateK := k % length
	current = head
	last.Next = head
	newHead := current
	for i := 0; i < length-rotateK-1; i++ {
		current = current.Next
	}
	newHead = current.Next
	current.Next = nil
	return newHead
}
