package main

import (
	"math"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func sum(l *ListNode) int64 {
	answer := int64(0)
	exp := 0
	for node := l; node != nil; node = node.Next {
		answer += int64(node.Val * int(math.Pow10(exp)))
		exp++
	}
	return answer
}

func convert(num int64) *ListNode {
	arg := num
	result := &ListNode{}
	current := result
	for {
		current.Val = int(arg % 10)
		arg = arg / 10
		if arg == 0 {
			return result
		}
		current.Next = &ListNode{}
		current = current.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carryOver := 0
	current1 := l1
	current2 := l2
	result := &ListNode{}
	resultCurrent := result
	for {
		val1 := 0
		if current1 != nil {
			val1 = current1.Val
		}
		val2 := 0
		if current2 != nil {
			val2 = current2.Val
		}
		val := val1 + val2 + carryOver
		resultCurrent.Val = val % 10
		carryOver = val / 10
		if current1 != nil {
			current1 = current1.Next
		}
		if current2 != nil {
			current2 = current2.Next
		}
		if current1 == nil && current2 == nil && carryOver == 0 {
			return result
		}
		resultCurrent.Next = &ListNode{}
		resultCurrent = resultCurrent.Next
	}
}

func TestConvert(t *testing.T) {
	arg := int64(342)
	t.Run("should be 2, 4, 3", func(t *testing.T) {
		answer := convert(arg)
		if answer.Val != 2 {
			t.Errorf("Not 2. actual: %v", answer)
		}
		if answer.Next.Val != 4 {
			t.Errorf("Not 4. actual: %v", answer)
		}
		if answer.Next.Next.Val != 3 {
			t.Errorf("Not 3, actual: %v", answer)
		}
		if answer.Next.Next.Next != nil {
			t.Errorf("Not nil, actual: %v", answer)
		}
	})
}


func TestSum(t *testing.T) {
	arg := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,				
			},
		},
	}
	t.Run("should be 342", func(t *testing.T) {
		answer := sum(arg)
		if answer != 342 {
			t.Errorf("actual: %v", answer)
		}
	})
	arg2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	t.Run("should be 101", func(t *testing.T) {
		answer := sum(arg2)
		if answer != 101 {
			t.Errorf("actual: %v", answer)
		}
	})
}

func TestAddTwoNumbers(t *testing.T) {

	arg1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	arg2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	t.Run("name", func(t *testing.T) {
		answer := addTwoNumbers(arg1, arg2)
		if answer.Val != 7 {
			t.Error("Not 7")
		}
		if answer.Next.Val != 0 {
			t.Error("Not 0")
		}
		if answer.Next.Next.Val != 8 {
			t.Error("Not 8")
		}
		if answer.Next.Next.Next != nil {
			t.Error("Not nil")
		}
	})

	arg1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},	
			},
		},
	}

	arg2 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
		},
	}

	t.Run("name2", func(t *testing.T) {
		answer := addTwoNumbers(arg1, arg2)
		if answer.Val != 8 {
			t.Errorf("acutal: %v", answer.Val)
		}
		if answer.Next.Val != 9 {
			t.Errorf("acutal: %v", answer.Next.Val)
		}
		if answer.Next.Next.Val != 0 {
			t.Errorf("acutal: %v", answer.Next.Next.Val)
		}
		if answer.Next.Next.Next.Val != 0 {
			t.Errorf("acutal: %v", answer.Next.Next.Next.Val)
		}
		if answer.Next.Next.Next.Next.Val != 1 {
			t.Errorf("acutal: %v", answer.Next.Next.Next.Next.Val)
		}
		if answer.Next.Next.Next.Next.Next != nil {
			t.Errorf("acutal: %v", answer)
		}
	})
}
