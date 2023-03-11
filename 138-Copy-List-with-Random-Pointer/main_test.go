package main

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {

	testcases := []struct {
		arg  *Node
		want *Node
	}{
		{
			arg:  nil,
			want: nil,
		},
		{
			arg: (func() *Node {
				n1 := &Node{
					Val: 7,
				}
				n2 := &Node{
					Val: 13,
				}
				n3 := &Node{
					Val: 11,
				}
				n4 := &Node{
					Val: 10,
				}
				n5 := &Node{
					Val: 1,
				}
				n1.Next = n2
				n2.Next = n3
				n3.Next = n4
				n4.Next = n5
				n2.Random = n1
				n3.Random = n5
				n4.Random = n3
				n5.Random = n1
				return n1
			}()),
			want: (func() *Node {
				n1 := &Node{
					Val: 7,
				}
				n2 := &Node{
					Val: 13,
				}
				n3 := &Node{
					Val: 11,
				}
				n4 := &Node{
					Val: 10,
				}
				n5 := &Node{
					Val: 1,
				}
				n1.Next = n2
				n2.Next = n3
				n3.Next = n4
				n4.Next = n5
				n2.Random = n1
				n3.Random = n5
				n4.Random = n3
				n5.Random = n1
				return n1
			}()),
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := copyRandomList(tc.arg)
			if !isEqual(actual, tc.want) {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}

func isEqual(a *Node, b *Node) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	for {
		if a == nil || b == nil {
			return a == nil && b == nil
		}
		if a.Val != b.Val {
			return false
		}
		if a.Random == nil || b.Random == nil {
			if a.Random != nil || b.Random != nil {
				return false
			}
		} else {
			if a.Random.Val != b.Random.Val {
				return false
			}
		}
		a = a.Next
		b = b.Next
	}
}
