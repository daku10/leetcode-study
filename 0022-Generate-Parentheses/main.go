package main

var memos = make(map[int][]string)

var memo = make(map[string]bool)

func generateParenthesis(n int) []string {
	if n == 1 {
		memos[1] = []string{"()"}
		return []string{"()"}
	}

	if memos[n] != nil {
		return memos[n]
	}

	result := make([]string, 0)
	for i := 1; i < n; i++ {
		tmp := generateParenthesis(n - i)
		lenTmp := len(tmp)
		tmp2 := generateParenthesis(i)
		lenTmp2 := len(tmp2)
		for j := 0; j < lenTmp; j++ {
			if i == 1 {
				result = append(result, "(" + tmp[j] + ")")
			}
			for k := 0;k < lenTmp2; k++ {
				candidate := tmp2[k] + tmp[j]
				if !memo[candidate] {
					memo[candidate] = true
					result = append(result, candidate)
				}
			}
		}
	}
	memos[n] = result
	return result
}
