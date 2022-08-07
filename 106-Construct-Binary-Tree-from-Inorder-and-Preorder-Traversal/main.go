package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	leftInorder, rightInorder := separateInorder(inorder, rootVal)
	leftPostorder, rightPostorder := separatePostorder(postorder, leftInorder)
	return &TreeNode{
		Val:   rootVal,
		Right: buildTree(rightInorder, rightPostorder),
		Left:  buildTree(leftInorder, leftPostorder),
	}
}

func separatePostorder(postorder []int, leftInorder []int) ([]int, []int) {
	for i := len(postorder) - 1; i >= 0; i-- {
		target := postorder[i]
		for j := 0; j < len(leftInorder); j++ {
			if target == leftInorder[j] {
				if i == len(postorder)-1 {
					return postorder, nil
				}
				return postorder[:i+1], postorder[i+1:]
			}
		}
	}
	return nil, postorder
}

func separateInorder(inorder []int, target int) ([]int, []int) {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			if i == 0 {
				return nil, inorder[1:]
			}
			if i == len(inorder)-1 {
				return inorder[:len(inorder)-1], nil
			}
			return inorder[:i], inorder[i+1:]
		}
	}
	panic("target must exist")
}
