package main

// ax1, ay1 left-bottom
// ax2, ay2 right-top
// bx1, by1 left-bottom
// bx2, by2 right-top
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	return ((ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - intersectionArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2))
}

func intersectionArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	w := min(ax2, bx2) - max(ax1, bx1)
	if w < 0 {
		return 0
	}

	h := min(ay2, by2) - max(ay1, by1)
	if h < 0 {
		return 0
	}

	return w * h
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
