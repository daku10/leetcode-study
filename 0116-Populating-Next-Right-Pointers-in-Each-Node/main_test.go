package main

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	testcases := []struct {
		arg  *Node
		want *Node
	}{
		{
			&Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val: 6,
					},
					Right: &Node{
						Val: 7,
					},
				},
			},
			func() *Node {
				one := &Node{
					Val: 1,
				}
				two := &Node{
					Val: 2,
				}
				three := &Node{
					Val: 3,
				}
				four := &Node{
					Val: 4,
				}
				five := &Node{
					Val: 5,
				}
				six := &Node{
					Val: 6,
				}
				seven := &Node{
					Val: 7,
				}
				one.Left = two
				one.Right = three
				two.Left = four
				two.Right = five
				three.Left = six
				three.Right = seven
				two.Next = three
				four.Next = five
				five.Next = six
				six.Next = seven
				return one
			}(),
		},
		{
			nil,
			nil,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := connect(tc.arg)
			if !cmpTree(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}

func cmpTree(x, y *Node) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if x.Next != nil && y.Next == nil {
		return false
	}
	if x.Next == nil && y.Next != nil {
		return false
	}
	if x.Val != y.Val {
		return false
	}
	if x.Next != nil && x.Next.Val != y.Next.Val {
		return false
	}
	return cmpTree(x.Left, y.Left) && cmpTree(x.Right, y.Right)
}
