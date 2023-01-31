package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	testcases := []struct {
		arg  [][]byte
		want [][]byte
	}{
		{
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			[][]byte{{'X'}},
			[][]byte{{'X'}},
		},
		{
			[][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
			[][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
			// output
			// [["O","O","O","O","X","X"],
			// ["O","O","O","O","O","O"],
			// ["O","X","O","X","O","O"],
			// ["O","X","O","X","X","O"],
			// ["O","X","O","X","O","O"],
			// ["O","X","O","O","O","O"]]
		},
		{
			[][]byte{
				{'O', 'X', 'O', 'O', 'X', 'X'},
				{'O', 'X', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'O', 'O'},
				{'X', 'O', 'X', 'X', 'X', 'X'},
				{'O', 'O', 'X', 'O', 'X', 'X'},
				{'X', 'X', 'O', 'O', 'O', 'O'},
			},
			[][]byte{
				{'O', 'X', 'O', 'O', 'X', 'X'},
				{'O', 'X', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'O', 'O'},
				{'X', 'O', 'X', 'X', 'X', 'X'},
				{'O', 'O', 'X', 'O', 'X', 'X'},
				{'X', 'X', 'O', 'O', 'O', 'O'},
			},
			// ["O","X","O","O","X","X"],
			// ["O","X","X","X","O","X"],
			// ["X","X","X","X","O","O"],
			// ["X","X","X","X","X","X"],
			// ["O","X","X","O","X","X"],
			// ["X","X","O","O","O","O"]
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			solve(tc.arg)
			if !reflect.DeepEqual(tc.want, tc.arg) {
				t.Errorf("want: %v actual: %v", tc.want, tc.arg)
			}
		})
	}
}
