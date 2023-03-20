package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pseudoHead := &ListNode{}
	pNode := pseudoHead
	originalHead := &ListNode{
		Next: head,
	}
	for originalHead.Next != nil {
		preCurrent := originalHead
		current := originalHead.Next
		min := current.Val
		minIndexPre := preCurrent
		for current != nil {
			if current.Val < min {
				min = current.Val
				minIndexPre = preCurrent
			}
			preCurrent = current
			current = current.Next
		}
		pNode.Next = minIndexPre.Next
		pNode = pNode.Next
		minIndexPre.Next = minIndexPre.Next.Next
	}
	return pseudoHead.Next
}
