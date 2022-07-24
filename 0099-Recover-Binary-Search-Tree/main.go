package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	original := convertToInorder(root, nil)
	sorted := make([]int, len(original))
	copy(sorted, original)
	sort.Ints(sorted)

	invalid0 := -1
	var invalid0Num, invalid1Num int

	for i := 0; i < len(original); i++ {
		if original[i] != sorted[i] {
			if invalid0 == -1 {
				invalid0 = i
				invalid0Num = original[i]
			} else {
				invalid1Num = original[i]
				break
			}
		}
	}
	traverse(root, invalid0Num, invalid1Num)
}

func traverse(node *TreeNode, invalid0Num int, invalid1Num int) {
	if node == nil {
		return
	}
	traverse(node.Left, invalid0Num, invalid1Num)
	if node.Val == invalid0Num {
		node.Val = invalid1Num
	} else if node.Val == invalid1Num {
		node.Val = invalid0Num
		return
	}
	traverse(node.Right, invalid0Num, invalid1Num)
}

func convertToInorder(root *TreeNode, arg []int) []int {
	if root == nil {
		return arg
	}
	var res []int
	res = convertToInorder(root.Left, arg)
	res = append(res, root.Val)
	res = convertToInorder(root.Right, res)
	return res
}
