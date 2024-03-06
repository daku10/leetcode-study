package main

func hIndex(citations []int) int {
	return binarySearch(citations, 0, len(citations)-1)
}

func binarySearch(citations []int, left int, right int) int {
	if right-left < 2 {
		switch right - left {
		case 1:
			if citations[left] >= len(citations)-left {
				return len(citations) - left
			}
			if citations[right] >= len(citations)-right {
				return len(citations) - right
			}
		case 0:
			if citations[left] >= len(citations)-left {
				return len(citations) - left
			}
		}
		return 0
	}
	mid := (left + right) / 2
	nums := len(citations) - mid
	if citations[mid] == nums {
		return nums
	}
	if citations[mid] > nums {
		return binarySearch(citations, left, mid)
	} else {
		return binarySearch(citations, mid, right)
	}
}
