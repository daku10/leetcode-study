package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	var nextQueue []*TreeNode
	isLeftDirection := true
	result := [][]int{}
	for {
		var res []int
		for i := 0; i < len(queue); i++ {
			q := queue[i]
			if isLeftDirection {
				if q.Left != nil {
					nextQueue = append(nextQueue, q.Left)
				}
				if q.Right != nil {
					nextQueue = append(nextQueue, q.Right)
				}
			} else {
				if q.Right != nil {
					nextQueue = append(nextQueue, q.Right)
				}
				if q.Left != nil {
					nextQueue = append(nextQueue, q.Left)
				}
			}
			res = append(res, q.Val)
		}
		if len(res) == 0 {
			break
		}
		result = append(result, res)
		isLeftDirection = !isLeftDirection
		queue = reverseDestructly(nextQueue)
		nextQueue = nil
	}
	return result
}

func reverseDestructly(arg []*TreeNode) []*TreeNode {
	for i, j := 0, len(arg)-1; i < j; {
		arg[i], arg[j] = arg[j], arg[i]
		i++
		j--
	}
	return arg
}
