package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	nodeMap := make(map[int][]*TreeNode)
	nodeMap[1] = []*TreeNode{
		{1, nil, nil},
	}
	for i := 2; i <= n; i++ {
		prevNodes := nodeMap[i-1]
		nodeMap[i] = generateNodes(prevNodes, i)
	}
	return nodeMap[n]
}

func generateNodes(prevNodes []*TreeNode, n int) []*TreeNode {
	result := make([]*TreeNode, 0)
	for _, pNode := range prevNodes {
		// top
		res := &TreeNode{
			Val:  n,
			Left: copyNode(pNode),
		}
		result = append(result, res)
		// right
		result = append(result, appendRightNode(pNode, n)...)

	}
	return result
}

func appendRightNode(prevNode *TreeNode, n int) []*TreeNode {
	result := make([]*TreeNode, 0)
	current := prevNode
	index := 0
	for current != nil {
		res := copyNode(prevNode)
		resCurrent := res
		for i := 0; i < index; i++ {
			resCurrent = resCurrent.Right
		}
		tmp := resCurrent.Right
		resCurrent.Right = &TreeNode{
			Val:  n,
			Left: tmp,
		}
		result = append(result, res)
		current = current.Right
		index++
	}

	return result
}

func copyNode(node *TreeNode) *TreeNode {
	var result *TreeNode = &TreeNode{
		Val: node.Val,
	}
	copyNodeDfs(node, result)
	return result
}

func copyNodeDfs(src *TreeNode, dst *TreeNode) {
	if src == nil {
		return
	}
	if src.Left != nil {
		dst.Left = &TreeNode{
			Val: src.Left.Val,
		}
	}
	if src.Right != nil {
		dst.Right = &TreeNode{
			Val: src.Right.Val,
		}
	}
	copyNodeDfs(src.Left, dst.Left)
	copyNodeDfs(src.Right, dst.Right)
}
