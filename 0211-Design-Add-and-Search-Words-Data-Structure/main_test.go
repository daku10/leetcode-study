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

func TestWordDictionary2(t *testing.T) {
	wd := Constructor()
	wd.AddWord("at")
	wd.AddWord("and")
	wd.AddWord("an")
	wd.AddWord("add")
	if wd.Search("a") {
		t.Errorf("should false")
	}
	if wd.Search(".at") {
		t.Errorf("should false")
	}
	wd.AddWord("bat")
	if !wd.Search(".at") {
		t.Errorf("should true")
	}
	if !wd.Search("an.") {
		t.Errorf("should true")
	}
	if wd.Search("a.d.") {
		t.Errorf("should false")
	}
	if wd.Search("b.") {
		t.Errorf("should false")
	}
	if !wd.Search("a.d") {
		t.Errorf("should true")
	}
	if wd.Search(".") {
		t.Errorf("should false")
	}
}
