package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	pseudoHead := &ListNode{
		Next: head,
		Val:  math.MinInt,
	}
	preCurr := head
	curr := head.Next
label:
	for curr != nil {
		preN := pseudoHead
		n := pseudoHead.Next
		for n != curr {
			if n.Val > curr.Val {
				preN.Next = curr
				tmp := curr.Next
				preCurr.Next = tmp
				curr.Next = n
				curr = tmp
				continue label
			}
			preN = n
			n = n.Next
		}
		preCurr = preCurr.Next
		curr = curr.Next
	}

	return pseudoHead.Next
}
