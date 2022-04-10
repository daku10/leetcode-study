package main

import (
	"strings"
)

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	length := len(paths)
	res := make([]string, length)
	index := 0
	for i := 0; i < length; i++ {
		dirName := paths[i]
		if dirName == "." {
			continue
		}
		if dirName == ".." {
			if index == 0 {
				continue
			}
			index--
			continue
		}
		if dirName == "" {
			continue
		}
		index++
		res[index] = dirName
	}

	if index == 0 {
		return "/"
	}

	result := strings.Builder{}
	result.WriteString("/")

	for i := 1; i < index; i++ {
		result.WriteString(res[i])
		result.WriteString("/")
	}
	result.WriteString(res[index])

	return result.String()
}
