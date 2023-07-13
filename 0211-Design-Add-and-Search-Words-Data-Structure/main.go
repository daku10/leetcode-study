package main

type WordDictionary struct {
	m        map[string]struct{}
	children [26]*Node
}

type Node struct {
	children [26]*Node
	hasEnd   bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		m: make(map[string]struct{}),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.m[word] = struct{}{}
	if this.children[word[0]-97] == nil {
		this.children[word[0]-97] = &Node{}
	}
	this.children[word[0]-97].AddWord(word[1:])
}

func (n *Node) AddWord(word string) {
	if word == "" {
		n.hasEnd = true
		return
	}
	if n.children[word[0]-97] == nil {
		n.children[word[0]-97] = &Node{}
	}
	n.children[word[0]-97].AddWord(word[1:])
}

func (this *WordDictionary) Search(word string) bool {
	if _, ok := this.m[word]; ok {
		return true
	}
	w := word[0]
	if w != '.' {
		if this.children[w-97] != nil {
			return this.children[w-97].Search(word[1:])
		}
		return false
	}
	if w == '.' {
		for _, c := range this.children {
			if c != nil && c.Search(word[1:]) {
				return true
			}
		}
	}
	return false
}

func (n *Node) Search(word string) bool {
	if word == "" {
		return n.hasEnd
	}
	if word[0] != '.' {
		if n.children[word[0]-97] == nil {
			return false
		}
		return n.children[word[0]-97].Search(word[1:])
	}
	for _, c := range n.children {
		if c != nil && c.Search(word[1:]) {
			return true
		}
	}
	return false
}
