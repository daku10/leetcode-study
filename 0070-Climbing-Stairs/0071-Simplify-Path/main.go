package main

import (
	"strings"
)

type Node struct {
	Val  string
	Prev *Node
}

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	root := &Node{
		Val:  "",
		Prev: nil,
	}
	current := root
	length := len(paths)
	for i := 0; i < length; i++ {
		dirName := paths[i]
		if dirName == "." {
			continue
		}
		if dirName == ".." {
			if current.Prev == nil {
				continue
			}
			current = current.Prev
			continue
		}
		if dirName == "" {
			continue
		}
		node := &Node{
			Val:  dirName,
			Prev: current,
		}
		current = node
	}

	if current.Prev == nil {
		return "/"
	}

	result := current.Val
	current = current.Prev
	for current != nil {
		result = current.Val + "/" + result
		current = current.Prev
	}

	return result
}
