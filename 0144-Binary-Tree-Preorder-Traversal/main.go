package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}
	for root != nil {
		if root.Left == nil {
			result = append(result, root.Val)
			root = root.Right
		} else {
			last := root.Left
			for last.Right != nil {
				last = last.Right
			}
			last.Right = root.Right
			result = append(result, root.Val)
			root = root.Left
		}
	}
	return result
}
