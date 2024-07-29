package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	if duration == 0 {
		return 0
	}
	var result int
	for i := 0; i < len(timeSeries)-1; i++ {
		result += min(duration, timeSeries[i+1]-timeSeries[i])
	}
	return result + duration
}
