package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type result struct {
	hasP bool
	hasQ bool
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	memo := make(map[*TreeNode]result)
	search(root, p, memo, true)
	search(root, q, memo, false)
	node := root
	for node != nil {
		if node.Left == nil {
			if v, ok := memo[node.Right]; ok {
				if v.hasP && v.hasQ {
					node = node.Right
					continue
				}
			}
			return node
		}
		if node.Right == nil {
			if v, ok := memo[node.Left]; ok {
				if v.hasP && v.hasQ {
					node = node.Left
					continue
				}
			}
			return node
		}
		l := memo[node.Left]
		r := memo[node.Right]
		if (l.hasP && !l.hasQ && !r.hasP && r.hasQ) || (!l.hasP && l.hasQ && r.hasP && !r.hasQ) {
			return node
		}
		if (l.hasP && !l.hasQ && !r.hasP && !r.hasQ) || (!l.hasP && l.hasQ && !r.hasP && !r.hasQ) {
			return node
		}
		if (!l.hasP && !l.hasQ && r.hasP && !r.hasQ) || (!l.hasP && !l.hasQ && !r.hasP && r.hasQ) {
			return node
		}
		if l.hasP && l.hasQ {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}

func search(node *TreeNode, target *TreeNode, memo map[*TreeNode]result, isSearcningP bool) result {
	if node == target {
		var res result
		if v, ok := memo[node]; ok {
			if isSearcningP {
				res = result{
					hasP: true,
					hasQ: v.hasQ,
				}
			} else {
				res = result{
					hasP: v.hasP,
					hasQ: true,
				}
			}
		} else {
			res = result{
				hasP: isSearcningP,
				hasQ: !isSearcningP,
			}
		}
		memo[node] = res
		return res
	}
	var l result
	var r result
	if node.Left != nil {
		l = search(node.Left, target, memo, isSearcningP)
	}
	if node.Right != nil {
		r = search(node.Right, target, memo, isSearcningP)
	}
	var resu result
	if v, ok := memo[node]; ok {
		resu = result{
			hasP: l.hasP || r.hasP || v.hasP,
			hasQ: l.hasQ || r.hasQ || v.hasQ,
		}
	} else {
		resu = result{
			hasP: l.hasP || r.hasP,
			hasQ: l.hasQ || r.hasQ,
		}
	}
	memo[node] = resu
	return resu
}
