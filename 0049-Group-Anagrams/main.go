package main

import (
	"reflect"
	"sort"
	"strings"
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
	if len(l) != len(r) {
		return false
	}
	lSplit := strings.Split(l, "")
	sort.Strings(lSplit)
	rSplit := strings.Split(r, "")
	sort.Strings(rSplit)
	return reflect.DeepEqual(lSplit, rSplit)
}
