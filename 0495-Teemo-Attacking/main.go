package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	var result int
	for i := 0; i < len(timeSeries)-1; i++ {
		t := timeSeries[i+1] - timeSeries[i]
		if t >= duration {
			result += duration
		} else {
			result += t
		}
	}
	return result + duration
}
