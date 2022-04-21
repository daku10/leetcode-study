package main

import (
	"reflect"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  func inorderTraversal(root *TreeNode) []int {

// }

func TestTraversal(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want []int
	}{
		{
			&TreeNode{
				1, nil, &TreeNode{
					2, &TreeNode{
						3, nil, nil,
					}, nil,
				},
			},
			[]int{1, 3, 2},
		},
		{
			nil,
			[]int{},
		},
		{
			&TreeNode{1, nil, nil},
			[]int{1},
		},
		{
			&TreeNode{
				1,
				&TreeNode{
					2,
					&TreeNode{
						4,
						nil, nil,
					},
					&TreeNode{
						5, nil, nil,
					},
				},
				&TreeNode{
					3, nil, nil,
				},
			},
			[]int{4, 2, 5, 1, 3},
		},
	}

	for _, tc := range testcases {
		actual := inorderTraversal(tc.arg)
		if !reflect.DeepEqual(tc.want, actual) {
			t.Errorf("expected: %v, actual: %v", tc.want, actual)
		}
	}
}
