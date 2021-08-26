package main

// Definition for singly-linked list.
type ListNode struct {
	Val int
    Next *ListNode
}
 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultCursor := result
	
	l1Current := l1
	l2Current := l2

	for {
		if l1Current == nil && l2Current == nil {
			return resultCursor.Next
		} else if l1Current == nil {
			result.Next = &ListNode{}
			result.Next.Val = l2Current.Val
			result = result.Next
			l2Current = l2Current.Next
		} else if l2Current == nil {
			result.Next = &ListNode{}
			result.Next.Val = l1Current.Val
			result = result.Next
			l1Current = l1Current.Next
		} else {
			if l1Current.Val < l2Current.Val {
				result.Next = &ListNode{}
				result.Next.Val = l1Current.Val
				result = result.Next
				l1Current = l1Current.Next	
			} else {
				result.Next = &ListNode{}
				result.Next.Val = l2Current.Val
				result = result.Next
				l2Current = l2Current.Next
			}
		}
	}
}