package main

import "testing"

func TestRangeSum(t *testing.T) {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	if na.SumRange(0, 2) != 1 {
		t.Fatalf("got: %v, want: 1", na.SumRange(0, 2))
	}
	if na.SumRange(2, 5) != -1 {
		t.Fatalf("got: %v, want: -1", na.SumRange(2, 5))
	}
	if na.SumRange(0, 5) != -3 {
		t.Fatalf("got: %v, want: -3", na.SumRange(0, 5))
	}
}
