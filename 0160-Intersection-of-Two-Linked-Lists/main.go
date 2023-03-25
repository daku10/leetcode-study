package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{})
	a := headA
	for a != nil {
		visited[a] = struct{}{}
		a = a.Next
	}
	b := headB
	for b != nil {
		if _, ok := visited[b]; ok {
			return b
		}
		b = b.Next
	}
	return nil
}
