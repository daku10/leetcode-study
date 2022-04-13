package main

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	testcases := []struct {
		argBoard [][]byte
		argWord  string
		want     bool
	}{
		{
			// {"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}
			argBoard: [][]byte{{65, 66, 67, 69}, {83, 70, 67, 83}, {65, 68, 69, 69}},
			argWord:  "ABCCED",
			want:     true,
		},
		{
			argBoard: [][]byte{{65, 66, 67, 69}, {83, 70, 67, 83}, {65, 68, 69, 69}},
			argWord:  "SEE",
			want:     true,
		},
		{
			argBoard: [][]byte{{65, 66, 67, 69}, {83, 70, 67, 83}, {65, 68, 69, 69}},
			argWord:  "ABCB",
			want:     false,
		},
		{
			argBoard: [][]byte{{65}, {65}},
			argWord:  "AAA",
			want:     false,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argBoard, tc.argWord), func(t *testing.T) {
			actual := exist(tc.argBoard, tc.argWord)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
