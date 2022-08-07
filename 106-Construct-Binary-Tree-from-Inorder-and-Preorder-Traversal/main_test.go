package main

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testcases := []struct {
		argInorder   []int
		argPostorder []int
		want         *TreeNode
	}{
		{
			argInorder:   []int{9, 3, 15, 20, 7},
			argPostorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			argInorder:   []int{-1},
			argPostorder: []int{-1},
			want: &TreeNode{
				Val: -1,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := buildTree(tc.argInorder, tc.argPostorder)
			if !cmpTree(got, tc.want) {
				t.Errorf("got: %+#v want: %+#v", got, tc.want)
			}
		})
	}
}

func cmpTree(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if x.Val != y.Val {
		return false
	}
	return cmpTree(x.Left, y.Left) && cmpTree(x.Left, y.Left)
}
