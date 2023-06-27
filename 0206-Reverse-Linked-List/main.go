package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	pseudoHead := &ListNode{}
	node := head
	for node != nil {
		tmp := node.Next
		node.Next = pseudoHead.Next
		pseudoHead.Next = node
		node = tmp
	}
	return pseudoHead.Next
}
