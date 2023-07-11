package main

import "testing"

func TestWordDictionary(t *testing.T) {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	if wd.Search("pad") {
		t.Errorf("should false")
	}
	if !wd.Search("bad") {
		t.Errorf("should true")
	}
	if !wd.Search(".ad") {
		t.Errorf("should true")
	}
	if !wd.Search("b..") {
		t.Errorf("should true")
	}
}
