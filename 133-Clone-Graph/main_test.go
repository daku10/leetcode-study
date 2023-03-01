package main

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {

	testcases := []struct {
		arg *Node
	}{
		{
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
				one.Neighbors = []*Node{two, four}
				two.Neighbors = []*Node{one, three}
				three.Neighbors = []*Node{two, four}
				four.Neighbors = []*Node{one, three}
				return one
			}(),
		},
		{
			&Node{
				Val: 1,
			},
		},
		{
			nil,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := cloneGraph(tc.arg)
			aVisited := make(map[*Node]struct{})
			bVisited := make(map[*Node]struct{})
			if !isEqual(actual, tc.arg, aVisited, bVisited) {
				t.Errorf("failed")
			}
		})
	}
}

func isEqual(a *Node, b *Node, aVisited map[*Node]struct{}, bVisited map[*Node]struct{}) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	_, aok := aVisited[a]
	_, bok := bVisited[b]
	if aok || bok {
		return aok && bok
	}
	if a.Val != b.Val {
		return false
	}
	if len(a.Neighbors) != len(b.Neighbors) {
		return false
	}
	aVisited[a] = struct{}{}
	bVisited[b] = struct{}{}

	for i := 0; i < len(a.Neighbors); i++ {
		if !isEqual(a.Neighbors[i], b.Neighbors[i], aVisited, bVisited) {
			return false
		}
	}
	return true
}
