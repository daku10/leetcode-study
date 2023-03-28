package main

func twoSum(numbers []int, target int) []int {
	for i, n := range numbers {
		s := target - n
		res := binarySearch(numbers, s, i, len(numbers)-1, i)
		if res != -1 {
			return []int{i + 1, res + 1}
		}
	}
	return nil
}

func binarySearch(nums []int, target int, start int, end int, ignore int) int {
	if start == end {
		if nums[start] == target && start != ignore {
			return start
		}
		return -1
	}
	if end-start == 1 {
		if nums[start] == target && start != ignore {
			return start
		}
		if nums[end] == target && end != ignore {
			return end
		}
		return -1
	}
	mid := (start + end) / 2
	numMid := nums[mid]
	if numMid == target {
		if mid != ignore {
			return mid
		}
		if nums[mid+1] == ignore {
			return mid + 1
		}
	}
	if numMid < target {
		return binarySearch(nums, target, mid, end, ignore)
	}
	return binarySearch(nums, target, start, mid, ignore)
}
