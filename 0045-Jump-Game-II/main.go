package main

func jump(nums []int) int {

	indexDistanceMap := make(map[int]int)
	indexDistanceMap[0] = 0

	length := len(nums)

	innerJump(nums, 0, length, indexDistanceMap)

	return indexDistanceMap[length - 1]
}

func innerJump(nums []int, index int, length int, indexDistanceMap map[int]int) {

	n := nums[index]
	distance := indexDistanceMap[index]
	
	if n == 0 {
		return
	}

	for i := 1; i <= n; i++ {
		next := index + i
		if next >= length {
			return
		}
		if v, ok := indexDistanceMap[next]; ok {
			if v > distance + 1 {
				indexDistanceMap[next] = distance + 1
			}
		} else {
			indexDistanceMap[next] = distance + 1
		}
		innerJump(nums, next, length, indexDistanceMap)
	}
}
