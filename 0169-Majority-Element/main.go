package main

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		v := m[n]
		if v+1 >= (len(nums)+1)/2 {
			return n
		}
		m[n] = v + 1
	}
	return 0
}
