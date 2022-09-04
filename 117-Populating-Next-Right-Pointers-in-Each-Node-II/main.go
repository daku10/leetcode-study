package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	var nextQueue []*Node
	for len(queue) != 0 {
		var nextN *Node
		for _, n := range queue {
			if n.Right != nil {
				nextQueue = append(nextQueue, n.Right)
			}
			if n.Left != nil {
				nextQueue = append(nextQueue, n.Left)
			}
			n.Next = nextN
			nextN = n
		}
		queue = nextQueue
		nextQueue = nil
	}
	return root
}
