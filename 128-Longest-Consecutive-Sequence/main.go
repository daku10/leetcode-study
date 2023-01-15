package main

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		m[n] = struct{}{}
	}

	result := 0

	for k := range m {
		if _, ok := m[k-1]; !ok {
			current := 1
			target := k
			for {
				if _, ok := m[target+1]; ok {
					current += 1
					target += 1
				} else {
					break
				}
			}
			if current > result {
				result = current
			}
		}
	}
	return result
}
