package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	result := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if q.Left != nil && q.Left.Left == nil && q.Left.Right == nil {
			result += q.Left.Val
		} else if q.Left != nil {
			queue = append(queue, q.Left)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}

	return result
}
