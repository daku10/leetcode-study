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
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			next := root.Next
			for next != nil {
				if next.Left != nil {
					root.Left.Next = next.Left
					break
				}
				if next.Right != nil {
					root.Left.Next = next.Right
					break
				}
				next = next.Next
			}
		}
	}
	if root.Right != nil {
		next := root.Next
		for next != nil {
			if next.Left != nil {
				root.Right.Next = next.Left
				break
			}
			if next.Right != nil {
				root.Right.Next = next.Right
				break
			}
			next = next.Next
		}
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
