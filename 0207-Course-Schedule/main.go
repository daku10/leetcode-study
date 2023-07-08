package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		inDegree[p[0]]++
	}

	var candidates []int
	for index, deg := range inDegree {
		if deg == 0 {
			candidates = append(candidates, index)
		}
	}
	numVisited := 0
	for len(candidates) != 0 {
		// take last elemenet from candidates, therefore this is not topological sort
		node := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		numVisited++
		ad := adj[node]
		for _, a := range ad {
			inDegree[a]--
			if inDegree[a] == 0 {
				candidates = append(candidates, a)
			}
		}
	}
	return numVisited == numCourses
}
