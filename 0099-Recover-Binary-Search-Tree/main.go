package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Object struct {
	Pre       *TreeNode
	Target    *TreeNode
	Candidate *TreeNode
}

func recoverTree(root *TreeNode) {
	var isFirstFound bool = false
	object := Object{}
	traverse(root, &isFirstFound, &object)
	if isFirstFound {
		tmp := object.Candidate.Val
		object.Candidate.Val = object.Target.Val
		object.Target.Val = tmp
	}
}

func traverse(node *TreeNode, isFirstFound *bool, object *Object) {
	if node.Left != nil {
		traverse(node.Left, isFirstFound, object)
	}
	if object.Pre != nil && object.Pre.Val > node.Val {
		if !*isFirstFound {
			*isFirstFound = true
			object.Target = object.Pre
			object.Candidate = node
		} else {
			tmp := node.Val
			node.Val = object.Target.Val
			object.Target.Val = tmp
			*isFirstFound = false
			return
		}
	}
	object.Pre = node
	if node.Right != nil {
		traverse(node.Right, isFirstFound, object)
	}
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
