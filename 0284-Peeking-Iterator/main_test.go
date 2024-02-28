package main

import "testing"

func TestIterator(t *testing.T) {
	i := &Iterator{
		Nums: []int{1, 2, 3},
	}
	got := i.next()
	if got != 1 {
		t.Errorf("got: %v, want: 1", got)
	}
	got = i.next()
	if got != 2 {
		t.Errorf("got: %v, want: 2", got)
	}
	got = i.next()
	if got != 3 {
		t.Errorf("got: %v, want: 3", got)
	}
	if i.hasNext() {
		t.Errorf("got: %v, want: false", i.hasNext())
	}
}

func TestPeekingIterator(t *testing.T) {
	p := Constructor(&Iterator{
		Nums: []int{1, 2, 3},
	})
	got := p.next()
	if got != 1 {
		t.Errorf("got: %v, want: 1", got)
	}
	got = p.peek()
	if got != 2 {
		t.Errorf("got: %v, want: 2", got)
	}
	got = p.next()
	if got != 2 {
		t.Errorf("got: %v, want: 2", got)
	}
	if !p.hasNext() {
		t.Errorf("got: false, want: true")
	}
	got = p.next()
	if got != 3 {
		t.Errorf("got: %v, want: 3", got)
	}
	if p.hasNext() {
		t.Errorf("got: true, want: false")
	}
}

func TestPeekingIteratorEmpty(t *testing.T) {
	p := Constructor(&Iterator{
		Nums: []int{},
	})
	if p.hasNext() {
		t.Errorf("got: true, want: false")
	}
}
