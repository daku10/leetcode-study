package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	var queue []*TreeNode = []*TreeNode{root}
	var nextQueue []*TreeNode
	for {
		var level []int
		for i := 0; i < len(queue); i++ {
			current := queue[i]
			if current.Left != nil {
				nextQueue = append(nextQueue, current.Left)
			}
			if current.Right != nil {
				nextQueue = append(nextQueue, current.Right)
			}
			level = append(level, current.Val)
		}
		if len(level) == 0 {
			break
		}
		queue = nextQueue
		nextQueue = nil
		result = append(result, level)
	}

	return result
}
