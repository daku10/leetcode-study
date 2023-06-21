package main

func isHappy(n int) bool {
	visited := make(map[int]struct{})
	current := n
	for {
		if current == 1 {
			return true
		}
		if _, ok := visited[current]; ok {
			return false
		}
		visited[current] = struct{}{}
		current = calc(current)
	}
}

func calc(n int) int {
	result := 0
	denom := 1000000000
	current := n
	for denom != 0 {
		q := current / denom
		result += (q * q)
		current -= (q * denom)
		denom /= 10
	}
	return result
}
