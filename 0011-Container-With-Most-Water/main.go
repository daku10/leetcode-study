package main

func maxArea(height []int) int {
	length := len(height)
	memo := make(map[int]bool, length)
	for i := 0; i < length; i++ {
		memo[i] = false
	}

	result := 0

	for i := 0; i < length; i++ {
		if memo[i] {
			continue
		}
		currentHeight := height[i]
		for j := i + 1; j < length; j++ {
			currentResult := 0
			if currentHeight > height[j] {
				memo[j] = true
				currentResult = height[j] * (j - i)
			} else if currentHeight == height[j] {
				memo[j] = true
				currentResult = currentHeight * (j - i)
			} else {
				currentResult = currentHeight * (j - i)
			}
			if currentResult > result {
				result = currentResult
			}
		}
	}

	return result
}
