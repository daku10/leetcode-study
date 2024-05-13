package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	testcases := []struct {
		turnedOn int
		want     []string
	}{
		{
			1,
			[]string{
				"0:01",
				"0:02",
				"0:04",
				"0:08",
				"0:16",
				"0:32",
				"1:00",
				"2:00",
				"4:00",
				"8:00",
			},
		},
		{
			9,
			nil,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := readBinaryWatch(tc.turnedOn)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
