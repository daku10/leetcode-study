package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return cloneRecursive(node, visited)
}

func cloneRecursive(node *Node, visited map[*Node]*Node) *Node {
	if cloned, ok := visited[node]; ok {
		return cloned
	}
	result := &Node{
		Val: node.Val,
	}
	visited[node] = result
	var neighbors []*Node
	for i := 0; i < len(node.Neighbors); i++ {
		neighbors = append(neighbors, cloneRecursive(node.Neighbors[i], visited))
	}
	result.Neighbors = neighbors
	return result
}
