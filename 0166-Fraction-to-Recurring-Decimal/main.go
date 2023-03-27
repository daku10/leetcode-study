package main

import (
	"fmt"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	quotient := numerator / denominator
	reminder := numerator % denominator
	if reminder == 0 {
		return fmt.Sprint(quotient)
	}
	if reminder < 0 {
		reminder *= -1
	}

	var result strings.Builder
	if quotient == 0 && ((numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0)) {
		result.WriteString("-")
	}
	result.WriteString(fmt.Sprintf("%d.", quotient))

	var decimal strings.Builder
	current := reminder * 10

	loopCount := 0
	loopMap := make(map[int]int)

	for current != 0 {
		if v, ok := loopMap[current]; ok {
			// loop Found
			d := decimal.String()
			result.WriteString(d[:v])
			result.WriteString("(")
			result.WriteString(d[v:])
			result.WriteString(")")
			return result.String()
		} else {
			loopMap[current] = loopCount
		}
		if current == denominator {
			return result.String()
		}
		if current < denominator {
			decimal.WriteString("0")
		} else {
			q := current / denominator
			if q < 0 {
				decimal.WriteString(fmt.Sprint(-1 * q))
			} else {
				decimal.WriteString(fmt.Sprint(q))
			}
			current -= (q * denominator)
		}
		current *= 10
		loopCount++
	}
	result.WriteString(decimal.String())
	return result.String()
}
