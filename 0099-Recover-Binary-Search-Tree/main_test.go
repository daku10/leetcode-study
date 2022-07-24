package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRecoverBinarySearchTree(t *testing.T) {

	testcases := []struct {
		arg  *TreeNode
		want *TreeNode
	}{
		{
			arg: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			recoverTree(tc.arg)
			if !reflect.DeepEqual(convertToInorder(tc.want, nil), convertToInorder(tc.arg, nil)) {
				t.Errorf("failed want: %v arg: %v", tc.want, tc.arg)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want []int
	}{
		{
			&TreeNode{
				Val: 1,
			},
			[]int{1},
		},
		{
			&TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			[]int{1, 2, 3},
		},
		{
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			[]int{1, 2, 3},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := convertToInorder(tc.arg, nil)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want=%v got=%v", tc.want, got)
			}
		})
	}
}
