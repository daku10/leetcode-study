package main

func searchMatrix(matrix [][]int, target int) bool {
	return binarySearchForMatrix(matrix, 0, len(matrix)*len(matrix[0]), target)
}

func binarySearchForMatrix(matrix [][]int, left int, right int, target int) bool {
	width := len(matrix[0])
	height := len(matrix)
	if right-left <= 1 {
		if matrix[left/width][left%width] == target {
			return true
		}
		if width*height > right && matrix[right/width][right%width] == target {
			return true
		}
		return false
	}
	middle := (left + right) / 2
	if matrix[middle/width][middle%width] == target {
		return true
	}
	if matrix[middle/width][middle%width] < target {
		return binarySearchForMatrix(matrix, middle, right, target)
	}
	return binarySearchForMatrix(matrix, left, middle, target)
}
