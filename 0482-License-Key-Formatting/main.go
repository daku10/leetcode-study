package main

import "strings"

func licenseKeyFormatting(s string, k int) string {
	ss := strings.ReplaceAll(strings.ToUpper(s), "-", "")
	if ss == "" {
		return ""
	}
	var result strings.Builder
	result.Grow(len(ss) + len(ss)/k)
	// write head
	f := len(ss) % k
	if f == 0 {
		f = k
	}
	result.WriteString(ss[:f])
	for i := f; i < len(ss); i += k {
		result.WriteByte('-')
		result.WriteString(ss[i : i+k])
	}
	return result.String()
}
