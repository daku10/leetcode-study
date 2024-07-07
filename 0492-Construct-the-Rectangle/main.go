package main

import "math"

func constructRectangle(area int) []int {
	sq := math.Sqrt(float64(area))
	sqi := int(math.Ceil(sq))
	for i := sqi; i <= area; i++ {
		if area%i == 0 {
			return []int{i, area / i}
		}
	}
	return nil
}
