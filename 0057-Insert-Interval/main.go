package main

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	newStart := newInterval[0]
	newEnd := newInterval[1]

	isOverlapping := false
	isNewIntervalMerged := false
	length := len(intervals)

	if length == 1 {
		iStart := intervals[0][0]
		iEnd := intervals[0][1]
		if isOverlap(intervals[0], newInterval) {
			start := iStart
			end := iEnd
			if newStart < iStart {
				start = newStart
			}
			if newEnd > iEnd {
				end = newEnd
			}
			return [][]int{{start, end}}
		} else {
			if iStart < newStart {
				return [][]int{{iStart, iEnd}, {newStart, newEnd}}
			} else {
				return [][]int{{newStart, newEnd}, {iStart, iEnd}}
			}
		}
	}

	var current []int

	for i := 0; i < length; i++ {
		interval := intervals[i]
		if isOverlapping {
			if isOverlap(interval, current) {
				start := interval[0]
				end := interval[1]
				if current[0] < start {
					start = current[0]
				}
				if current[1] > end {
					end = current[1]
				}
				current[0] = start
				current[1] = end
			} else {
				result = append(result, current)
				result = append(result, interval)
				isOverlapping = false
			}
		} else {
			if isOverlap(interval, newInterval) {
				start := interval[0]
				end := interval[1]
				if newInterval[0] < interval[0] {
					start = newInterval[0]
				}
				if newInterval[1] > interval[1] {
					end = newInterval[1]
				}
				current = []int{start, end}
				isOverlapping = true
				isNewIntervalMerged = true
			} else {
				if newInterval[0] < interval[0] && !isNewIntervalMerged {
					result = append(result, newInterval)
					result = append(result, interval)
					isNewIntervalMerged = true
				} else {
					result = append(result, interval)
				}
			}
		}
	}

	if isOverlapping {
		result = append(result, current)
	}

	if !isNewIntervalMerged {
		result = append(result, newInterval)
	}

	return result
}

func isOverlap(i1, i2 []int) bool {
	return i1[0] <= i2[1] && i1[1] >= i2[0]
}
