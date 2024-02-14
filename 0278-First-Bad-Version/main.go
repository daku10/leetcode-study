package main

var BadVersion int

func isBadVersion(version int) bool {
	return version >= BadVersion
}

func firstBadVersion(n int) int {
	return firstBadVersionRecursive(1, n)
}

func firstBadVersionRecursive(low int, high int) int {
	mid := (high + low) / 2
	if isBadVersion(mid) {
		if low == mid {
			return mid
		}
		if !isBadVersion(mid - 1) {
			return mid
		}
		return firstBadVersionRecursive(low, mid-1)
	} else {
		if isBadVersion(mid + 1) {
			return mid + 1
		}
		return firstBadVersionRecursive(mid+1, high)
	}
}
