package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	current := head
	dummy := &ListNode{
		Val:  1000000,
		Next: head,
	}
	lastNode := dummy
	duplicatedVal := 100000

	for current != nil && current.Next != nil {
		if current.Val != current.Next.Val {
			lastNode.Next = current
			lastNode = lastNode.Next
			current = current.Next
		} else {
			duplicatedVal = current.Val
			for current != nil && current.Val == duplicatedVal {
				current = current.Next
			}
		}
	}

	if current == nil {
		lastNode.Next = nil
	} else if duplicatedVal != current.Val {
		lastNode.Next = current
	}

	return dummy.Next
}
