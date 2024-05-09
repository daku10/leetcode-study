package main

import (
	"math"
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
var answer = 0

func guess(num int) int {
	if num == answer {
		return 0
	}
	if num > answer {
		return -1
	}
	return 1
}

func guessNumber(n int) int {
	var low int64 = 1
	var high int64 = math.MaxInt32
	for {
		current := (low + high) / 2
		c := int(current)
		n := guess(c)
		if n == 0 {
			return c
		}
		if guess(int(high)) == 0 {
			return int(high)
		}
		if n == -1 {
			high = current
		} else {
			low = current
		}
	}
}
