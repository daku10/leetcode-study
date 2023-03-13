package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	visited := make(map[*ListNode]struct{})
	n := head
	for n != nil {
		if _, ok := visited[n]; ok {
			return true
		}
		visited[n] = struct{}{}
		n = n.Next
	}
	return false
}
