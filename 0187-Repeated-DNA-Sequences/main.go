package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	var result []string
	memo := make(map[string]struct{})
	resultMemo := make(map[string]struct{})
	for i := 0; i <= len(s)-10; i++ {
		target := s[i : i+10]
		if _, ok := memo[target]; ok {
			if _, ok := resultMemo[target]; ok {
				continue
			}
			result = append(result, target)
			resultMemo[target] = struct{}{}
		} else {
			memo[target] = struct{}{}
		}
	}
	return result
}
