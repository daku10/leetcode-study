package main

import "testing"

func TestBinarySearchIterator(t *testing.T) {
	tree := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	iter := Constructor(tree)
	if v := iter.Next(); v != 3 {
		t.Fatalf("got: %v", v)
	}
	if v := iter.Next(); v != 7 {
		t.Fatalf("got: %v", v)
	}
	if v := iter.HasNext(); v != true {
		t.Fatalf("got: false")
	}
	if v := iter.Next(); v != 9 {
		t.Fatalf("got: %v", v)
	}
	if v := iter.HasNext(); v != true {
		t.Fatalf("got: false")
	}
	if v := iter.Next(); v != 15 {
		t.Fatalf("got: %v", v)
	}
	if v := iter.HasNext(); v != true {
		t.Fatalf("got: false")
	}
	if v := iter.Next(); v != 20 {
		t.Fatalf("got: %v", v)
	}
	if v := iter.HasNext(); v != false {
		t.Fatalf("got: true")
	}
}
