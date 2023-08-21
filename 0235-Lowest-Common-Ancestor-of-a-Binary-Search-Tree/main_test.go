package main

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	type arg struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}

	testcases := []arg{
		(func() arg {
			p := &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			}
			q := &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			}
			root := &TreeNode{
				Val:   6,
				Left:  p,
				Right: q,
			}
			return arg{
				root: root,
				p:    p,
				q:    q,
				want: root,
			}
		})(),
		(func() arg {
			q := &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			}
			p := &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
				Right: q,
			}

			root := &TreeNode{
				Val:  6,
				Left: p,
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			}
			return arg{
				root: root,
				p:    p,
				q:    q,
				want: p,
			}
		})(),
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := lowestCommonAncestor(tc.root, tc.p, tc.q)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
