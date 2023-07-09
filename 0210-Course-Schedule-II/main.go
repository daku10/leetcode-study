package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		inDegree[pre[0]]++
	}
	var result []int
	var numVisited int
	var queue []int
	for index, d := range inDegree {
		if d == 0 {
			queue = append(queue, index)
		}
	}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		numVisited++
		for _, ad := range adj[node] {
			inDegree[ad]--
			if inDegree[ad] == 0 {
				queue = append(queue, ad)
			}
		}
	}
	if numVisited == numCourses {
		return result
	}
	return []int{}
}
