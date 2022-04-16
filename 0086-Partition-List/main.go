package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	current := head
	var beforeCurrent *ListNode = nil
	var beforePartiionNode *ListNode = &ListNode{
		Next: head,
	}
	result := head
	doneFirstSwap := false
	for current != nil {
		if current.Val < x {
			if !doneFirstSwap {
				doneFirstSwap = true
				result = current
			}
			if beforeCurrent == nil {
				beforeCurrent = current
				beforePartiionNode = current
				current = current.Next
			} else if beforePartiionNode.Next != current {
				beforeCurrent.Next = current.Next
				current.Next = beforePartiionNode.Next
				beforePartiionNode.Next = current
				beforePartiionNode = beforePartiionNode.Next
				current = beforeCurrent.Next
			} else {
				beforeCurrent = current
				beforePartiionNode = beforePartiionNode.Next
				current = current.Next
			}
		} else {
			beforeCurrent = current
			current = current.Next
		}
	}
	return result
}
