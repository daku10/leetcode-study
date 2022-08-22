package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	testcases := []struct {
		argNode   *TreeNode
		argTarget int
		want      [][]int
	}{
		{
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			22,
			[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			5,
			[][]int{},
		},
		{
			nil,
			0,
			[][]int{},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := pathSum(tc.argNode, tc.argTarget)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v argNode: %v argTarget: %v", got, tc.want, tc.argNode, tc.argTarget)
			}
		})
	}
}
