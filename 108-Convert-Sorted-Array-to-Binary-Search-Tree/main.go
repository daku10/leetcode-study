package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	mid := length / 2
	var left *TreeNode
	if mid > 0 {
		left = sortedArrayToBST(nums[:mid])
	}
	right := sortedArrayToBST(nums[mid+1:])
	return &TreeNode{
		Val:   nums[mid],
		Left:  left,
		Right: right,
	}
}
