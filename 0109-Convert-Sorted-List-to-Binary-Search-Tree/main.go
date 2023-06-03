package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	node := head
	length := 1
	for node.Next != nil {
		node = node.Next
		length++
	}
	return sortedListToBSTSub(head, node, length)
}

func sortedListToBSTSub(head *ListNode, end *ListNode, length int) *TreeNode {
	if head == nil {
		return nil
	}
	if head == end {
		return &TreeNode{
			Val: head.Val,
		}
	}
	node := head
	for node != end {
		node = node.Next
	}
	mid := length / 2
	midPrevNode := head
	for i := 0; i < mid-1; i++ {
		midPrevNode = midPrevNode.Next
	}
	var left *TreeNode
	if midPrevNode != head {
		left = sortedListToBSTSub(head, midPrevNode, mid-1)
	} else {
		left = &TreeNode{
			Val: head.Val,
		}
	}
	var right *TreeNode
	if midPrevNode.Next != nil && midPrevNode.Next != end {
		right = sortedListToBSTSub(midPrevNode.Next.Next, end, length-mid-1)
	}
	return &TreeNode{
		Val:   midPrevNode.Next.Val,
		Left:  left,
		Right: right,
	}
}
