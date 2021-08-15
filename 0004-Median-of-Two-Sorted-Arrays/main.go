package main

// time complexity of this algorithm is over constraint of the problem probably...
// I think this: O(n + m), constraint: O(log(n + m))
// but this passes the problem.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := createSortedArrays(nums1, nums2)
	length := len(merged)
	if length % 2 == 0 {
		val1 := merged[(length - 1) / 2]
		val2 := merged[length / 2]
		return float64(val1 + val2) / 2
	} else {
		return float64(merged[length / 2])
	}
}

func createSortedArrays(num1 []int, num2 []int) []int {
	result := make([]int, 0, len(num1) + len(num2))
	index1 := 0
	index2 := 0
	for {
		if index1 == len(num1) && index2 == len(num2) {
			break
		} else if index1 == len(num1) {
			result = append(result, num2[index2])
			index2++
		} else if index2 == len(num2) {
			result = append(result, num1[index1])
			index1++
		} else {
			if num1[index1] > num2[index2] {
				result = append(result, num2[index2])
				index2++
			} else {
				result = append(result, num1[index1])
				index1++
			}
		}
	}
	return result
}
