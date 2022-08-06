package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	preorder = preorder[1:]
	leftInorder, rightInorder := separateInorder(inorder, rootVal)
	leftPreorder, rightPreorder := separatePreorder(preorder, rightInorder)

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(leftPreorder, leftInorder),
		Right: buildTree(rightPreorder, rightInorder),
	}
}

func separatePreorder(preorder []int, rightInorder []int) ([]int, []int) {
	if len(rightInorder) == 0 {
		return preorder, nil
	}
	for i := 0; i < len(preorder); i++ {
		target := preorder[i]
		for j := 0; j < len(rightInorder); j++ {
			if target == rightInorder[j] {
				return preorder[:i], preorder[i:]
			}
		}
	}
	panic("must be unreachable")
}

// left, right
func separateInorder(inorder []int, target int) ([]int, []int) {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			if i == 0 {
				return nil, inorder[1:]
			} else if i == len(inorder) {
				return inorder[:len(inorder)-1], nil
			}
			return inorder[:i], inorder[i+1:]
		}
	}
	panic("target must exist")
}
