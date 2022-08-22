package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	current := []int{}
	pathSumSub(root, targetSum, &current, &result)
	return result
}

func pathSumSub(node *TreeNode, targetSum int, current *[]int, result *[][]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if node.Val == targetSum {
			res := make([]int, len(*current))
			copy(res, *current)
			res = append(res, node.Val)
			*result = append(*result, res)
		}
		return
	}
	*current = append(*current, node.Val)
	pathSumSub(node.Left, targetSum-node.Val, current, result)
	pathSumSub(node.Right, targetSum-node.Val, current, result)
	*current = (*current)[:len(*current)-1]
}
