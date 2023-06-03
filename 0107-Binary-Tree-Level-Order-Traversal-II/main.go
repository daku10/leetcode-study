package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	nextQueue := []*TreeNode{}
	for {
		var res []int
		for _, q := range queue {
			res = append(res, q.Val)
			if q.Left != nil {
				nextQueue = append(nextQueue, q.Left)
			}
			if q.Right != nil {
				nextQueue = append(nextQueue, q.Right)
			}
		}
		if len(res) == 0 {
			break
		}
		result = append(result, res)
		res = nil
		queue = nextQueue
		nextQueue = nil
	}
	// reverse(result)
	return result
}

func reverse(arg [][]int) {
	for i, j := 0, len(arg)-1; i < j; {
		arg[i], arg[j] = arg[j], arg[i]
		i++
		j--
	}
}
