package main

type Node struct {
	Next []*Node
}

func NewNode(v byte) *Node {
	return &Node{
		Next: make([]*Node, 26),
	}
}

type Trie struct {
	dict  map[string]struct{}
	Nodes []*Node
}

func Constructor() Trie {
	return Trie{
		make(map[string]struct{}),
		make([]*Node, 26),
	}
}

func (this *Trie) Insert(word string) {
	this.dict[word] = struct{}{}
	p := word[0]
	node := this.Nodes[p-97]
	if node == nil {
		node = NewNode(p)
		this.Nodes[p-97] = node
	}
	if len(word) == 1 {
		return
	}
	node.Insert(word[1:])
}

func (n *Node) Insert(word string) {
	p := word[0]
	node := n.Next[p-97]
	if node == nil {
		node = NewNode(p)
		n.Next[p-97] = node
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
	n := this.Nodes[p-97]
	if n == nil {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	return n.StartsWith(prefix[1:])
}

func (n *Node) StartsWith(prefix string) bool {
	p := prefix[0]
	next := n.Next[p-97]
	if next == nil {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	return next.StartsWith(prefix[1:])
}
