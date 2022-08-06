package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var preorderIndex int
	return buildTreeSub(preorder, &preorderIndex, inorder)
}

func buildTreeSub(preorder []int, preorderIndex *int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(preorder) == *preorderIndex {
		return nil
	}
	rootVal := preorder[*preorderIndex]
	*preorderIndex++
	leftInorder, rightInorder := separateInorder(inorder, rootVal)

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTreeSub(preorder, preorderIndex, leftInorder),
		Right: buildTreeSub(preorder, preorderIndex, rightInorder),
	}
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
