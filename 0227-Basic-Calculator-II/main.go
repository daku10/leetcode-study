package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Val  string
	Prev *Node
	Next *Node
}

func calculate(s string) int {
	pseudo := &Node{}
	node := pseudo
	var str strings.Builder
	muldivCount := 0
	for _, r := range s {
		switch r {
		case ' ':
			continue
		case '+', '-', '*', '/':
			n := &Node{
				Val:  str.String(),
				Prev: node,
			}
			node.Next = n
			node = n
			str.Reset()
			n2 := &Node{
				Val:  string(r),
				Prev: node,
			}
			node.Next = n2
			node = n2
			if r == '*' || r == '/' {
				muldivCount++
			}
		default:
			str.WriteByte(byte(r))
		}
	}
	node.Next = &Node{
		Val:  str.String(),
		Prev: node,
	}

	node = pseudo

	for node != nil {
		switch node.Val {
		case "*", "/":
			pre := toInt(node.Prev.Val)
			nex := toInt(node.Next.Val)
			newNode := &Node{}
			if node.Val == "*" {
				newNode.Val = strconv.Itoa(pre * nex)
			} else {
				newNode.Val = strconv.Itoa(pre / nex)
			}
			n1 := node.Prev.Prev
			n2 := node.Next.Next
			n1.Next = newNode
			newNode.Prev = n1
			newNode.Next = n2
			if n2 != nil {
				n2.Prev = newNode
			}
			node = newNode
		}
		node = node.Next
	}

	node = pseudo
	for node != nil {
		switch node.Val {
		case "+", "-":
			pre := toInt(node.Prev.Val)
			nex := toInt(node.Next.Val)
			newNode := &Node{}
			if node.Val == "+" {
				newNode.Val = strconv.Itoa(pre + nex)
			} else {
				newNode.Val = strconv.Itoa(pre - nex)
			}
			n1 := node.Prev.Prev
			n2 := node.Next.Next
			n1.Next = newNode
			newNode.Prev = n1
			newNode.Next = n2
			if n2 != nil {
				n2.Prev = newNode
			}
			node = newNode
		}
		node = node.Next
	}

	return toInt(pseudo.Next.Val)
}

func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
