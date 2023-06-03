package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want []int
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
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			[]int{1, 3, 4},
		},
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
			},
			[]int{1, 3},
		},
		{
			nil,
			nil,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := rightSideView(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
