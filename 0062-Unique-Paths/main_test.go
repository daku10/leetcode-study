package main

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	testcases := []struct {
		argM int
		argN int
		want int
	}{
		{
			argM: 3,
			argN: 7,
			want: 28,
		},
		{
			argM: 3,
			argN: 2,
			want: 3,
		},
		{
			argM: 16,
			argN: 16,
			want: 155117520,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argM, tc.argN), func(t *testing.T) {
			actual := uniquePaths(tc.argM, tc.argN)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
