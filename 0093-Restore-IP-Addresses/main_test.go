package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestRestore(t *testing.T) {
	testcases := []struct {
		arg  string
		want []string
	}{
		{
			"25525511135",
			[]string{"255.255.11.135", "255.255.111.35"},
		},
		{
			"0000",
			[]string{"0.0.0.0"},
		},
		{
			"101023",
			[]string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			"255255255255",
			[]string{"255.255.255.255"},
		},
		{
			"333333333333333",
			[]string{},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.arg), func(t *testing.T) {
			actual := restoreIpAddresses(tc.arg)
			sort.Strings(tc.want)
			sort.Strings(actual)
			if !reflect.DeepEqual(tc.want, actual) {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
