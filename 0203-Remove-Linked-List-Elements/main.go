package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	pseudoHead := &ListNode{
		Val:  0,
		Next: head,
	}
	n := pseudoHead
	for n.Next != nil {
		if n.Next.Val == val {
			tmp := n.Next.Next
			n.Next.Next = nil
			n.Next = tmp
			continue
		}
		n = n.Next
	}
	return pseudoHead.Next
}
