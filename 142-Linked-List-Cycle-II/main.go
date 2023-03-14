package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	visited := make(map[*ListNode]struct{})
	n := head
	for n != nil {
		if _, ok := visited[n]; ok {
			return n
		}
		visited[n] = struct{}{}
		n = n.Next
	}
	return nil
}
