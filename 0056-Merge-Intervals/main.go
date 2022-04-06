package main

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	length := len(intervals)
	var tmp []int
	memo := make(map[int]bool)
	isOverlapCalc := false

	for i := 0; i < length; i++ {
		if memo[i] {
			continue
		}
		tmp = []int{intervals[i][0], intervals[i][1]}
		j := i + 1
		for {
			if j >= length {
				if isOverlapCalc {
					j = i + 1
					isOverlapCalc = false
				} else {
					break
				}
			}
			if memo[j] {
				j++
				continue
			}
			compare := intervals[j]
			iStart := tmp[0]
			iEnd := tmp[1]
			jStart := compare[0]
			jEnd := compare[1]
			// check overlap
			if (iStart <= jStart && iEnd >= jStart) || (jStart <= iStart && jEnd >= iStart) {
				end := jEnd
				start := iStart
				if iEnd > jEnd {
					end = iEnd
				}
				if jStart < iStart {
					start = jStart
				}
				tmp = []int{start, end}
				memo[j] = true
				isOverlapCalc = true

			}
			j++
		}
		result = append(result, tmp)
	}
	return result
}
