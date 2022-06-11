package main

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	testcases := []struct {
		argS1 string
		argS2 string
		argS3 string
		want  bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"aabcc", "dbbca", "aadbcbbcac", true},
		{
			"bcbccabcccbcbbbcbbacaaccccacbaccabaccbabccbabcaabbbccbbbaa",
			"ccbccaaccabacaabccaaccbabcbbaacacaccaacbacbbccccbac",
			"bccbcccabbccaccaccacbacbacbabbcbccbaaccbbaacbcbaacbacbaccaaccabcaccacaacbacbacccbbabcccccbababcaabcbbcccbbbaa",
			true,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.argS1, tc.argS2, tc.argS3), func(t *testing.T) {
			actual := isInterleave(tc.argS1, tc.argS2, tc.argS3)
			if tc.want != actual {
				t.Errorf("expected: %v, actual: %v", tc.want, actual)
			}
		})
	}
}
