package main

func searchMatrix(matrix [][]int, target int) bool {
	height := len(matrix)
	width := len(matrix[0])

	targetHeight, found := binarySearchForHeightIndex(matrix, 0, height-1, target)
	if !found {
		return false
	}
	return binarySearch(matrix[targetHeight], 0, width-1, target)
}

func binarySearchForHeightIndex(matrix [][]int, min int, max int, target int) (int, bool) {
	minValue := matrix[min][0]
	maxValue := matrix[max][0]
	if minValue == target {
		return min, true
	}
	if maxValue == target {
		return max, true
	}
	if minValue > target {
		return min, false
	}
	// search last column
	if maxValue < target {
		return max, true
	}
	if max-min <= 1 {
		return min, true
	}
	middle := (max + min) / 2
	middleValue := matrix[middle][0]
	if middleValue == target {
		return middle, true
	}
	if middleValue > target {
		return binarySearchForHeightIndex(matrix, min, middle-1, target)
	}
	return binarySearchForHeightIndex(matrix, middle, max, target)
}

func binarySearch(arr []int, left int, right int, target int) bool {
	if right-left <= 1 {
		if arr[right] == target || arr[left] == target {
			return true
		}
		return false
	}
	middle := (left + right) / 2
	if arr[middle] == target {
		return true
	}
	if arr[middle] < target {
		return binarySearch(arr, middle, right, target)
	}
	return binarySearch(arr, left, middle, target)
}
