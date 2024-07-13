package main

import "math"

func constructRectangle(area int) []int {
	sq := int(math.Sqrt(float64(area)))
	for i := sq; i > 0; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return nil
}
