package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var vals []int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		vals = append(vals, q.Val)
		if q.Left != nil {
			queue = append(queue, q.Left)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}
	sort.Ints(vals)

	return vals[k-1]
}
