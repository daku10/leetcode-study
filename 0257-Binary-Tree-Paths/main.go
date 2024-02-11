package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	current := []string{}
	dfs(root, &current, &result)
	return result
}

func dfs(node *TreeNode, current *[]string, result *[]string) {
	v := strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		s := strings.Join((*current), "->")
		if len(s) == 0 {
			s = v
		} else {
			s += "->" + v
		}
		*result = append(*result, s)
		return
	}
	*current = append(*current, strconv.Itoa(node.Val))
	if node.Left != nil {
		dfs(node.Left, current, result)
	}
	if node.Right != nil {
		dfs(node.Right, current, result)
	}
	*current = (*current)[:len(*current)-1]
}
