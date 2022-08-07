package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderIndex := len(postorder) - 1
	return buildTreeSub(inorder, postorder, &postorderIndex)
}

func buildTreeSub(inorder []int, postorder []int, postorderIndex *int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if *postorderIndex == -1 {
		return nil
	}
	rootVal := postorder[*postorderIndex]
	*postorderIndex--
	leftInorder, rightInorder := separateInorder(inorder, rootVal)
	// must call Right first to decrement index properly
	return &TreeNode{
		Val:   rootVal,
		Right: buildTreeSub(rightInorder, postorder, postorderIndex),
		Left:  buildTreeSub(leftInorder, postorder, postorderIndex),
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
