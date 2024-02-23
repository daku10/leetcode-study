package main

import "testing"

func TestRangeSumQuery(t *testing.T) {
	m := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	got := m.SumRegion(2, 1, 4, 3)
	if got != 8 {
		t.Errorf("got: %v, want: 8", got)
	}
	got = m.SumRegion(1, 1, 2, 2)
	if got != 11 {
		t.Errorf("got: %v, want: 11", got)
	}
	got = m.SumRegion(1, 2, 2, 4)
	if got != 12 {
		t.Errorf("got: %v, want: 12", got)
	}
}
