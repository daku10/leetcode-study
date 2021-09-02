package main


type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var length int

	items := make(map[int]*ListNode)
	itemIndex := 0

	current := head
	for current != nil {
		items[itemIndex] = current
		itemIndex++		
		current = current.Next
	}
	length = itemIndex

	target := length - n
	if target == 0 {
		return head.Next
	}

	items[target - 1].Next = items[target + 1]

	return head
}
