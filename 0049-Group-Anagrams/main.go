package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	result := make([][]string, 0)

	length := len(strs)

	memo := make(map[string]bool)

	for i := 0; i < length; i++ {
		str := strs[i]

		if memo[str] {
			continue
		}

		res := []string{str}
		
		for j := i + 1; j < length; j++ {
			compareStr := strs[j]
			if isAnagram(str, compareStr) {
				memo[compareStr] = true
				res = append(res, compareStr)
			}
		}
		result = append(result, res)
	}

    return result
}

func isAnagram(l string, r string) bool {
	length := len(l)
	if length != len(r) {
		return false
	}
	lSplit := make([]byte, length)
	rSplit := make([]byte, length)
	for i := 0; i < length; i++ {
		lSplit[i] = l[i]
		rSplit[i] = r[i]
	}	

	sort.Slice(lSplit, func(i, j int) bool {
		return lSplit[i] < lSplit[j]
	});
	sort.Slice(rSplit, func(i, j int) bool {
		return rSplit[i] < rSplit[j]
	})

	for i := 0; i < length; i++ {
		if lSplit[i] != rSplit[i] {
			return false
		}
	}
	
	return true
}
