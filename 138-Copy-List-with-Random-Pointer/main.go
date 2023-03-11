package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	memo := make(map[*Node]*Node)
	h := head
	copiedHead := &Node{
		Val: h.Val,
	}
	ch := copiedHead
	memo[h] = copiedHead
	for h != nil {
		h = h.Next
		if h != nil {
			n := &Node{
				Val: h.Val,
			}
			memo[h] = n
			ch.Next = n
			ch = n
		}
	}
	h = head
	ch = copiedHead
	for h != nil {
		ch.Random = memo[h.Random]
		h = h.Next
		ch = ch.Next
	}

	return copiedHead
}
