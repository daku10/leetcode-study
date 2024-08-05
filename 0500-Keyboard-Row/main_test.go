package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	testcases := []struct {
		words []string
		want  []string
	}{
		{
			[]string{"Hello", "Alaska", "Dad", "Peace"},
			[]string{"Alaska", "Dad"},
		},
		{
			[]string{"omk"},
			nil,
		},
		{
			[]string{"adsdf", "sfd"},
			[]string{"adsdf", "sfd"},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findWords(tc.words)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
