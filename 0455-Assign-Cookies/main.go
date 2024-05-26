package main

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	var i, j int
	for i < len(g) && j < len(s) {
		if g[i] > s[j] {
			j++
			continue
		}
		i++
		j++
	}
	return i
}
