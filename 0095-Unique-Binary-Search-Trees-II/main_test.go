package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCopyNode(t *testing.T) {
	testcases := []struct {
		arg  *TreeNode
		want *TreeNode
	}{
		{
			&TreeNode{
				Val: 1,
			},
			&TreeNode{
				Val: 1,
			},
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
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
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
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
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
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := copyNode(tc.arg)
			if !reflect.DeepEqual(preorder(tc.arg), preorder(actual)) {
				t.Errorf("expected: %v, actual: %v", preorder(tc.want), preorder(actual))
			}
		})
	}
}

func TestGenerateTrees(t *testing.T) {
	testcases := []struct {
		arg  int
		want []*TreeNode
	}{
		{
			3,
			[]*TreeNode{
				{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, nil}}},
				{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, nil}},
				{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
				{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, nil},
				{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil},
			},
		},
		{
			1,
			[]*TreeNode{{1, nil, nil}},
		},
		{
			2,
			[]*TreeNode{
				{1, nil, &TreeNode{2, nil, nil}},
				{2, &TreeNode{1, nil, nil}, nil},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := generateTrees(tc.arg)
			actualInorder := make([][]int, 0)
			for _, ac := range actual {
				actualInorder = append(actualInorder, preorder(ac))
			}
			wantInorder := make([][]int, 0)
			for _, wc := range tc.want {
				wantInorder = append(wantInorder, preorder(wc))
			}
			sort.Slice(actualInorder, func(i, j int) bool {
				ac1 := actualInorder[i]
				ac2 := actualInorder[j]
				for k := 0; k < len(ac1); k++ {
					if ac1[k] < ac2[k] {
						return true
					} else if ac1[k] > ac2[k] {
						return false
					}
				}
				return false
			})

			sort.Slice(wantInorder, func(i, j int) bool {
				wc1 := wantInorder[i]
				wc2 := wantInorder[j]
				for k := 0; k < len(wc1); k++ {
					if wc1[k] < wc2[k] {
						return true
					} else if wc1[k] > wc2[k] {
						return false
					}
				}
				return false
			})

			if !reflect.DeepEqual(wantInorder, actualInorder) {
				t.Errorf("expected: %v, actual: %v", wantInorder, actualInorder)
			}
		})
	}
}

func preorder(node *TreeNode) []int {
	result := make([]int, 0)
	preorderDfs(node, &result)
	return result
}

func preorderDfs(node *TreeNode, result *[]int) {
	if node == nil {
		*result = append(*result, -1)
		return
	}
	*result = append(*result, node.Val)
	preorderDfs(node.Left, result)
	preorderDfs(node.Right, result)
}
