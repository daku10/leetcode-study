package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	memo := make(map[int]int)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		memo[q.Val]++
		if q.Left != nil {
			queue = append(queue, q.Left)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}

	return findMax(memo)
}

func findMax(m map[int]int) []int {
	var max int
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	var result []int
	for k, v := range m {
		if v == max {
			result = append(result, k)
		}
	}
	return result
}
