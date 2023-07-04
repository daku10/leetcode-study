package main

type Node struct {
	Val  int
	Next []*Node
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	memo := make(map[int]*Node)
	heads := make(map[*Node]struct{})

	for _, p := range prerequisites {
		after := p[0]
		before := p[1]
		var afterNode, beforeNode *Node
		if v, ok := memo[after]; ok {
			afterNode = v
		} else {
			afterNode = &Node{
				Val: after,
			}
			memo[after] = afterNode
		}
		if v, ok := memo[before]; ok {
			beforeNode = v
		} else {
			beforeNode = &Node{
				Val: before,
			}
			memo[before] = beforeNode
		}
		beforeNode.Next = append(beforeNode.Next, afterNode)
		delete(heads, afterNode)
		heads[beforeNode] = struct{}{}
	}

	for k := range heads {
		visited := make(map[*Node]struct{})
		okNodes := make(map[*Node]struct{})
		if hasCycle(k, visited, okNodes) {
			return false
		}
	}

	return true
}

func hasCycle(current *Node, visited map[*Node]struct{}, okNodes map[*Node]struct{}) bool {
	if _, ok := visited[current]; ok {
		return true
	}
	if _, ok := okNodes[current]; ok {
		return false
	}
	if current.Next == nil {
		return false
	}
	visited[current] = struct{}{}
	for _, n := range current.Next {
		if hasCycle(n, visited, okNodes) {
			return true
		}
	}
	delete(visited, current)
	okNodes[current] = struct{}{}
	return false
}
