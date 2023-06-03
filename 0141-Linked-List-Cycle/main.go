package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	n1 := head
	n2 := head.Next
	for {
		if n1 == n2 {
			return true
		}
		if n2 == nil {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
		if n2 == nil {
			return false
		}
		n2 = n2.Next
	}
}
