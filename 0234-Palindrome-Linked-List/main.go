package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return isPalindromeArr(arr)
}

func isPalindromeArr(x []int) bool {
	i := 0
	j := len(x) - 1
	for i <= j {
		if x[i] != x[j] {
			return false
		}
		i++
		j--
	}
	return true
}
