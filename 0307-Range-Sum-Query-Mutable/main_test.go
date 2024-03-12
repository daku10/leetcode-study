package main

import "testing"

func TestRangeSumQuery(t *testing.T) {
	numArr := Constructor([]int{1, 3, 5})
	got := numArr.SumRange(0, 2)
	if got != 9 {
		t.Errorf("got: %v, want: %v", got, 9)
	}
	numArr.Update(1, 2)
	got = numArr.SumRange(0, 2)
	if got != 8 {
		t.Errorf("got: %v, want: %v", got, 8)
	}
}
