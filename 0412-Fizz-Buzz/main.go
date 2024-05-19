package main

import "fmt"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		r := ""
		if i%3 == 0 {
			r += "Fizz"
		}
		if i%5 == 0 {
			r += "Buzz"
		}
		if r == "" {
			r = fmt.Sprint(i)
		}
		result[i-1] = r
	}
	return result
}
