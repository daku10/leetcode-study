package main

type ListNode struct {
     Val int
     Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	result := head.Next
	current := head
	for {
		if current == nil {
			break
		}
		if current.Next == nil {
			break
		}
		if current.Next.Next == nil {
			current.Next.Next = current
			current.Next = nil
			break;
		} else if current.Next.Next.Next == nil {
			tmp := current.Next.Next			
			current.Next.Next = current
			current.Next = tmp
			current = current.Next
		} else {
			tmp := current.Next.Next.Next
			next := current.Next.Next
			current.Next.Next = current
			current.Next = tmp
			current = next
		}

	}
	return result
}
