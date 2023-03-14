package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	fast := head
	slow := head
	for {
		if fast == nil {
			break
		}
		if fast.Next == nil {
			slow = slow.Next
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow points the first node of the second half
	// reverse the second half
	var pre *ListNode
	n := slow
	for n != nil {
		tmp := n.Next
		n.Next = pre
		pre = n
		n = tmp
	}
	// pre is the head of the second half
	// merge the first half list and the second half list
	n = head
	for n != nil {
		tmp := n.Next
		n.Next = pre
		n = tmp
		if pre == nil {
			return
		}
		tmp = pre.Next
		pre.Next = n
		pre = tmp
	}
}
