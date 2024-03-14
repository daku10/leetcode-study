package main

import (
	"fmt"
	"math"
)

type treeNode struct {
	Val       int
	Connected []*treeNode
}

func findMinHeightTrees(n int, edges [][]int) []int {
	nodes := createTree(n, edges)
	minHeight := math.MaxInt
	heights := make(map[int]int)
	// key = node_from
	memo := make(map[string]int)
	for _, node := range nodes {
		height := calcHeight(node, nil, memo)
		heights[node.Val] = height
		if height < minHeight {
			minHeight = height
		}
	}
	var result []int
	for k, v := range heights {
		if v == minHeight {
			result = append(result, k)
		}
	}
	return result
}

func calcHeight(node *treeNode, from *treeNode, memo map[string]int) int {
	if len(node.Connected) == 0 {
		return 0
	}
	if len(node.Connected) == 1 && node.Connected[0] == from {
		return 0
	}
	fromVal := -1
	if from != nil {
		fromVal = from.Val
	}
	key := fmt.Sprintf("%v_%v", node.Val, fromVal)
	if v, ok := memo[key]; ok {
		return v
	}
	var maxHeight int
	for _, con := range node.Connected {
		if con == from {
			continue
		}
		height := calcHeight(con, node, memo) + 1
		if maxHeight < height {
			maxHeight = height
		}
	}
	memo[key] = maxHeight
	return maxHeight
}

func createTree(n int, edges [][]int) []*treeNode {
	result := make([]*treeNode, n)
	for i := 0; i < n; i++ {
		t := &treeNode{
			Val: i,
		}
		result[i] = t
	}
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		e1 := edge[0]
		e2 := edge[1]
		result[e1].Connected = append(result[e1].Connected, result[e2])
		result[e2].Connected = append(result[e2].Connected, result[e1])
	}
	return result
}
