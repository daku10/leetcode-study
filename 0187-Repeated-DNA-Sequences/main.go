package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	var result []string
	memo := make(map[string]struct{})
	for i := 0; i <= len(s)-10; i++ {
		target := s[i : i+10]
		if _, ok := memo[target]; ok {
			continue
		}
		for j := i + 1; j <= len(s)-10; j++ {
			if target == s[j:j+10] {
				result = append(result, target)
				memo[target] = struct{}{}
				break
			}
		}
	}
	return result
}
