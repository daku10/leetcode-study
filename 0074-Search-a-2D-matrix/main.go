package main

func searchMatrix(matrix [][]int, target int) bool {
	height := len(matrix)
	width := len(matrix[0])
	for y := 0; y < height; y++ {
		if matrix[y][0] == target || matrix[y][width-1] == target {
			return true
		}
		if matrix[y][0] < target && matrix[y][width-1] > target {
			return binarySearch(matrix[y], 0, width-1, target)
		}
	}
	return false
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
