package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}
	var leftQueue []*TreeNode = []*TreeNode{root.Left}
	var rightQueue []*TreeNode = []*TreeNode{root.Right}
	var left *TreeNode
	var right *TreeNode

	for {
		if len(leftQueue) == 0 && len(rightQueue) == 0 {
			return true
		}
		if len(leftQueue) == 0 || len(rightQueue) == 0 {
			return false
		}
		left = leftQueue[0]
		right = rightQueue[0]
		leftQueue = leftQueue[1:]
		rightQueue = rightQueue[1:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		leftQueue = append(leftQueue, left.Left, left.Right)
		rightQueue = append(rightQueue, right.Right, right.Left)
	}

	return true
}

func print(arg []*int) {
	for _, a := range arg {
		if a == nil {
			fmt.Print("nil")
		} else {
			fmt.Print(*a)
		}
		fmt.Print(" ")
	}
}

func traverseInOrder(node *TreeNode, res []*int) []*int {
	if node == nil {
		return append(res, nil)
	}
	result := traverseInOrder(node.Left, res)
	result = append(result, &node.Val)
	return traverseInOrder(node.Right, result)
}

func traversePostOrder(node *TreeNode, res []*int) []*int {
	if node == nil {
		return append(res, nil)
	}
	result := traversePostOrder(node.Right, res)
	result = append(result, &node.Val)
	return traversePostOrder(node.Left, result)
}
