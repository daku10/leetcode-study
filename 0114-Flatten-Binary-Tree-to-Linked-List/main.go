package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	node := root
	preorderTraversal(node)
}

func preorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	if node.Left != nil {
		tmpRight := node.Right
		rightMostNode := node.Left
		for rightMostNode.Right != nil {
			rightMostNode = rightMostNode.Right
		}
		rightMostNode.Right = tmpRight
		node.Right = node.Left
		node.Left = nil
	}
	preorderTraversal(node.Right)
}
