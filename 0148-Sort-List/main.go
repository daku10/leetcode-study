package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Val < head.Next.Val {
			return head
		}
		tmp := head.Next
		head.Next = nil
		tmp.Next = head
		return tmp
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mergeList(sortList(head), sortList(mid))
}

func mergeList(a *ListNode, b *ListNode) *ListNode {
	pseudo := &ListNode{}
	p := pseudo
	for a != nil || b != nil {
		if a == nil {
			p.Next = b
			p = p.Next
			b = b.Next
		} else if b == nil {
			p.Next = a
			p = p.Next
			a = a.Next
		} else {
			if a.Val < b.Val {
				p.Next = a
				p = p.Next
				a = a.Next
			} else {
				p.Next = b
				p = p.Next
				b = b.Next
			}
		}
	}
	return pseudo.Next
}
