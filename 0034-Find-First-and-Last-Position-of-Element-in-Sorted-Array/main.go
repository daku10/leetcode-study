package main

func searchRange(nums []int, target int) []int {
    length := len(nums)
    targetIndex := innerSearchRange(nums, 0, length - 1, target)
    if targetIndex == -1 {
        return []int{-1, -1}
    } else if nums[targetIndex] != target {
        return []int{-1, -1}
    }

    lowestIndex := searchBoundaryIndex(nums, 0, targetIndex, target, true)
    highestIndex := searchBoundaryIndex(nums, targetIndex, length - 1, target, false)

    return []int{lowestIndex, highestIndex}
}

// this is just binary searchex
func innerSearchRange(nums[] int, l int, r int, target int) int {
    if (l < 0) || (r < 0) {
        return -1
    }
    m := (l + r) / 2
    left := nums[l]
    right := nums[r]
    middle := nums[m]

    if target == middle {
        return m
    }
    if target == left {
        return l
    }
    if target == right {
        return r
    }
    if (r - l) < 2  {
        return -1
    }
    if target > right || target < left {
        return -1
    }
    if target < middle {
        return innerSearchRange(nums, l, m - 1, target)
    } else {
        return innerSearchRange(nums, m + 1, r, target)
    }
}

// search boundary index
func searchBoundaryIndex(nums []int, l int, r int, target int, isLowest bool) int {
    if l == r {
        return l
    }
    if (r - l) < 2 {
        if isLowest {
            if nums[l] == target {
                return l
            }
            return r
        }
        if nums[r] == target {
            return r
        }
        return l
    }
    m := (l + r) / 2
    middle := nums[m]

    if isLowest {
        if middle < target {
            return searchBoundaryIndex(nums, m + 1, r, target, isLowest)
        } else {
            return searchBoundaryIndex(nums, l, m, target, isLowest)
        }
    }
    if middle > target {
        return searchBoundaryIndex(nums, l, m - 1, target, isLowest)
    } else {
        return searchBoundaryIndex(nums, m, r, target, isLowest)
    }
}
