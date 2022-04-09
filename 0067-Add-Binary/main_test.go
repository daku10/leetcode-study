package main

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	testcases := []struct {
		argA string
		argB string
		want string
	}{
		{
			argA: "11",
			argB: "1",
			want: "100",
		},
		{
			argA: "1010",
			argB: "1011",
			want: "10101",
		},
		{
			argA: "0",
			argB: "0",
			want: "0",
		},
		{
			argA: "0",
			argB: "1",
			want: "1",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argA, tc.argB), func(t *testing.T) {
			actual := addBinary(tc.argA, tc.argB)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
