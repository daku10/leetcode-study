package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	pInLeft := search(root.Left, p)
	qInLeft := search(root.Left, q)

	if pInLeft != qInLeft {
		return root
	}
	if pInLeft && qInLeft {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}

func search(root *TreeNode, target *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		q := queue[0]
		if q == target {
			return true
		}
		queue = queue[1:]
		if q.Left != nil {
			queue = append(queue, q.Left)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}
	return false
}
