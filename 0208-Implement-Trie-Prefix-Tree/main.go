package main

type Node struct {
	Val  byte
	Next map[byte]*Node
}

func NewNode(v byte) *Node {
	return &Node{
		Val:  v,
		Next: make(map[byte]*Node),
	}
}

type Trie struct {
	dict  map[string]struct{}
	Nodes map[byte]*Node
}

func Constructor() Trie {
	return Trie{
		make(map[string]struct{}),
		make(map[byte]*Node),
	}
}

func (this *Trie) Insert(word string) {
	this.dict[word] = struct{}{}
	p := word[0]
	node, ok := this.Nodes[p]
	if !ok {
		node = NewNode(p)
		this.Nodes[p] = node
	}
	if len(word) == 1 {
		return
	}
	node.Insert(word[1:])
}

func (n *Node) Insert(word string) {
	p := word[0]
	node, ok := n.Next[p]
	if !ok {
		node = NewNode(p)
		n.Next[p] = node
	}
	if len(word) == 1 {
		return
	}
	node.Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if _, ok := this.dict[word]; ok {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	p := prefix[0]
	n, ok := this.Nodes[p]
	if !ok {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	return n.StartsWith(prefix[1:])
}

func (n *Node) StartsWith(prefix string) bool {
	p := prefix[0]
	next, ok := n.Next[p]
	if !ok {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	return next.StartsWith(prefix[1:])
}
