package main

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("should true")
	}
	if trie.Search("app") {
		t.Errorf("should false")
	}
	if !trie.StartsWith("app") {
		t.Errorf("should true")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("should true")
	}
}
