package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var preorderIndex int
	inorderMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}
	return buildTreeSub(preorder, &preorderIndex, 0, len(inorder)-1, inorderMap)
}

func buildTreeSub(preorder []int, preorderIndex *int, left int, right int, inorderMap map[int]int) *TreeNode {
	if left > right {
		return nil
	}
	if len(preorder) == *preorderIndex {
		return nil
	}
	rootVal := preorder[*preorderIndex]
	*preorderIndex++
	pos := inorderMap[rootVal]

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTreeSub(preorder, preorderIndex, left, pos-1, inorderMap),
		Right: buildTreeSub(preorder, preorderIndex, pos+1, right, inorderMap),
	}
}
