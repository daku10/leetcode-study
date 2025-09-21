package main

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	globalMin := nums[0]
	globalMax := nums[0]
	for _, n := range nums {
		globalMax = max(globalMax, n)
		globalMin = min(globalMin, n)
	}
	if globalMax-globalMin < 2 {
		return globalMax - globalMin
	}

	// each bucket has at most average of gaps, result must exist between buckets.
	buckets := make([]*struct {
		l int
		h int
	}, len(nums))

	for _, n := range nums {
		i := (n - globalMin) * (len(nums) - 1) / (globalMax - globalMin)
		if buckets[i] == nil {
			buckets[i] = &struct {
				l int
				h int
			}{
				l: n,
				h: n,
			}
		} else {
			buckets[i].h = max(buckets[i].h, n)
			buckets[i].l = min(buckets[i].l, n)
		}
	}

	result := 0
	for i := 0; i < len(buckets); {
		if buckets[i] == nil {
			i++
			continue
		}
		for j := i + 1; j < len(buckets); j++ {
			if buckets[j] != nil {
				result = max(result, buckets[j].l-buckets[i].h)
				i = j
				continue
			}
		}
		return max(result, (buckets[i].h - buckets[i].l))
	}

	return result
}
