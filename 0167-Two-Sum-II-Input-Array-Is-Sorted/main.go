package main

func twoSum(numbers []int, target int) []int {
	for i, n := range numbers {
		s := target - n
		res := search(numbers, s, i)
		if res != -1 {
			return []int{i + 1, res + 1}
		}
	}
	return nil
}

func search(nums []int, s int, ignoreIndex int) int {
	for i, n := range nums {
		if i == ignoreIndex {
			continue
		}
		if n == s {
			return i
		}
		if n > s {
			return -1
		}
	}
	return -1
}
