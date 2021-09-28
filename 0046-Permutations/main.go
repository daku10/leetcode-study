package main

func permute(nums []int) [][]int {

	length := len(nums)

	result := make([][]int, 0)

	for i := 0; i < length; i++ {
		n := nums[i]
		rest := make([]int, 0)
		for j := 0; j < length; j++ {
			if n != nums[j] {
				rest = append(rest, nums[j])
			}
		}
		inner([]int{n}, rest, &result)
	}

    return result
}

func inner(current []int, rest []int, result *[][]int) {
	lenRest := len(rest)
	if lenRest == 0 {
		*result = append((*result), current)
		return
	}
	for i := 0; i < lenRest; i++ {
		r := rest[i]
		restRest := make([]int, 0)
		for j := 0; j < lenRest; j++ {
			if r != rest[j] {
				restRest = append(restRest, rest[j])
			}
		}
		newCurrent := make([]int, 0)
		newCurrent = append(newCurrent, current...)
		newCurrent = append(newCurrent, r)
		inner(newCurrent, restRest, result)
	}
}