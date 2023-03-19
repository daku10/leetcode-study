package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want []int
	}{
		{
			func() *TreeNode {
				n0 := &TreeNode{
					Val: 1,
				}
				n1 := &TreeNode{
					Val: 2,
				}
				n2 := &TreeNode{
					Val: 3,
				}
				n0.Right = n1
				n1.Left = n2
				return n0
			}(),
			[]int{1, 2, 3},
		},
		{
			nil,
			nil,
		},
		{
			&TreeNode{Val: 1},
			[]int{1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := preorderTraversal(tc.arg)
			if !reflect.DeepEqual(actual, tc.want) {
				t.Errorf("actual: %v want: %v", actual, tc.want)
			}
		})
	}
}
