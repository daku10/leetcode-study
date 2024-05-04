package main

import (
	"reflect"
	"testing"
)

func TestFlattenNestedListIterator(t *testing.T) {
	i := Constructor([]*NestedInteger{
		&NestedInteger{kind: List, l: []*NestedInteger{
			&NestedInteger{kind: Int, i: 1},
			&NestedInteger{kind: Int, i: 1},
		}},
		&NestedInteger{kind: Int, i: 2},
		&NestedInteger{kind: List, l: []*NestedInteger{
			&NestedInteger{kind: Int, i: 1},
			&NestedInteger{kind: Int, i: 1},
		}},
	})
	var res []int
	for i.HasNext() {
		res = append(res, i.Next())
	}
	want := []int{1, 1, 2, 1, 1}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("got: %v, want: %v", res, want)
	}

	i = Constructor([]*NestedInteger{
		&NestedInteger{kind: Int, i: 1},
		&NestedInteger{kind: List, l: []*NestedInteger{
			&NestedInteger{kind: Int, i: 4},
			&NestedInteger{kind: List, l: []*NestedInteger{
				&NestedInteger{kind: Int, i: 6},
			}},
		}},
	})
	res = []int{}
	for i.HasNext() {
		res = append(res, i.Next())
	}
	want = []int{1, 4, 6}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("got: %v, want: %v", res, want)
	}

}
