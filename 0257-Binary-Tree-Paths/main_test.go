package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want []string
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			[]string{
				"1->2->5",
				"1->3",
			},
		},
		{
			&TreeNode{
				Val: 1,
			},
			[]string{"1"},
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			[]string{
				"1->2->4",
				"1->2->5",
				"1->3",
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := binaryTreePaths(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
