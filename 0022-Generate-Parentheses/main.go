package main

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	innerGenerate(n, n, &result)
	return result
}

func innerGenerate(n int, target int, r *[]string) {
	if n < target && n != 1 {
		innerGenerate(n - 1, target, r)
	}
	
	tmp := ""
	for i := 0; i < n; i++ {
		tmp += "("
	}
	for i := 0; i < n; i++ {
		tmp += ")"
	}
	r = append(r)
}
