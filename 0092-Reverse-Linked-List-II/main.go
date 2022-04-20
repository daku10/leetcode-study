package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	preHead := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	preHead = pre
	pre = pre.Next
	current := pre.Next

	for i := 0; i < (right - left); i++ {
		tmp := current
		pre.Next = current.Next
		current = current.Next
		tmp.Next = preHead.Next
		preHead.Next = tmp
	}

	return dummy.Next
}
