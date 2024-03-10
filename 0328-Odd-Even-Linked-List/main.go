package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pseudoOddhead := &ListNode{}
	currentOddHead := pseudoOddhead
	pseudoEvenHead := &ListNode{}
	currentEvenHead := pseudoEvenHead
	currentOddHead.Next = head
	currentOddHead = currentOddHead.Next
	currentEvenHead.Next = head.Next
	currentEvenHead = currentEvenHead.Next
	for {
		if currentEvenHead.Next == nil || currentEvenHead.Next.Next == nil {
			if currentEvenHead.Next == nil {
				currentOddHead.Next = pseudoEvenHead.Next
				return pseudoOddhead.Next
			}
			currentOddHead.Next = currentEvenHead.Next
			currentOddHead.Next.Next = pseudoEvenHead.Next
			currentEvenHead.Next = nil
			return pseudoOddhead.Next
		}
		currentOddHead.Next = currentEvenHead.Next
		currentOddHead = currentOddHead.Next
		currentEvenHead.Next = currentOddHead.Next
		currentEvenHead = currentEvenHead.Next
	}
}
