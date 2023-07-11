package main

import "strings"

type WordDictionary struct {
	m map[string]struct{}
}

func Constructor() WordDictionary {
	return WordDictionary{
		m: make(map[string]struct{}),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.m[word] = struct{}{}
}

func (this *WordDictionary) Search(word string) bool {
	if _, ok := this.m[word]; ok {
		return true
	}
	if strings.ContainsRune(word, '.') {
		for i := 0; i < 26; i++ {
			alph := string([]byte{(byte)(97 + i)})
			if this.Search(strings.Replace(word, ".", alph, 1)) {
				return true
			}
		}
	}
	return false
}
