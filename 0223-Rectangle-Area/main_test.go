package main

import (
	"fmt"
	"testing"
)

func TestComputeArea(t *testing.T) {
	testcases := []struct {
		ax1  int
		ay1  int
		ax2  int
		ay2  int
		bx1  int
		by1  int
		bx2  int
		by2  int
		want int
	}{
		{
			-3,
			0,
			3,
			4,
			0,
			-1,
			9,
			2,
			45,
		},
		{
			-2,
			-2,
			2,
			2,
			-2,
			-2,
			2,
			2,
			16,
		},
		{
			-10,
			-10,
			0,
			0,
			-5,
			-5,
			5,
			-2,
			115,
		},
		{
			-10,
			-10,
			10,
			10,
			-5,
			-5,
			5,
			5,
			400,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := computeArea(tc.ax1, tc.ay1, tc.ax2, tc.ay2, tc.bx1, tc.by1, tc.bx2, tc.by2)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
