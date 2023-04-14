package main

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	testcases := []struct {
		num  uint32
		want uint32
	}{
		{
			43261596,
			964176192,
		},
		{
			4294967293,
			3221225471,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reverseBits(tc.num)
			if got != tc.want {
				t.Errorf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
