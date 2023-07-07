package main

import "strings"

type Trie struct {
	t []string
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	this.t = append(this.t, word)
}

func (this *Trie) Search(word string) bool {
	for _, node := range this.t {
		if word == node {
			return true
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, node := range this.t {
		if strings.HasPrefix(node, prefix) {
			return true
		}
	}
	return false
}
