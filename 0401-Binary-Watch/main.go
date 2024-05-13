package main

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	if turnedOn >= 9 {
		return nil
	}
	var result []string
	for i := 0; i <= turnedOn; i++ {
		hour := readHour(i)
		minute := readMinute(turnedOn - i)
		if hour != nil && minute != nil {
			for _, h := range hour {
				for _, m := range minute {
					result = append(result, h+":"+m)
				}
			}
		}
	}
	return result
}

func readHour(n int) []string {
	if n == 0 {
		return []string{"0"}
	}
	if n == 1 {
		return []string{"1", "2", "4", "8"}
	}
	if n == 2 {
		return []string{"3", "5", "6", "9", "10"}
	}
	if n == 3 {
		return []string{"7", "11"}
	}
	return nil
}

func readMinute(n int) []string {
	if n >= 6 {
		return nil
	}
	if n == 0 {
		return []string{"00"}
	}
	var result []string
	for i := 1; i < 60; i++ {
		if count(i) == n {
			result = append(result, fmt.Sprintf("%02d", i))
		}
	}
	return result
}

func count(n int) int {
	result := 0
	for i := 0; i < 6; i++ {
		if (n & (1 << i)) > 0 {
			result++
		}
	}
	return result
}
