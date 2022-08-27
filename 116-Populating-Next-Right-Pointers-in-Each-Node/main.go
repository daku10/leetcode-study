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
		var tmp *Node
		for _, n := range queue {
			if tmp != nil {
				n.Next = tmp
			}
			tmp = n
			if n.Right != nil {
				nextQueue = append(nextQueue, n.Right)
			}
			if n.Left != nil {
				nextQueue = append(nextQueue, n.Left)
			}
		}
		queue = nextQueue
		nextQueue = nil
	}

	return root
}
