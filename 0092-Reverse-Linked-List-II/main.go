package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	current := head
	reverseBeforeHead := &ListNode{
		Next: head,
	}
	reverseHead := &ListNode{}
	var reverseCurrent *ListNode = nil
	isSearching := false
	index := 1
	for current != nil {
		if index == left {
			isSearching = true
			reverseHead = current
		} else if index == right {
			reverseHead.Next = current.Next
			reverseBeforeHead.Next = current
			current.Next = reverseCurrent
			break
		}

		tmp := current
		current = current.Next
		if isSearching {
			tmp.Next = reverseCurrent
			reverseCurrent = tmp
		} else {
			reverseBeforeHead = reverseBeforeHead.Next
		}
		index++
	}
	if left == 1 {
		return current
	}
	return head
}
