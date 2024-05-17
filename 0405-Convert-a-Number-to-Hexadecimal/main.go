package main

import (
	"fmt"
	"strings"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		res := toHexCore(-1 * num)
		return strings.TrimLeft(increment(invert(res)), "0")
	}

	return strings.TrimLeft(toHexCore(num), "0")
}

func toHexCore(num int) string {
	st := 1 << 28
	current := num
	var result string
	for st != 0 {
		s := current / st
		result += toHexdecimal(s)
		current = current % st
		st /= 16
	}
	return result
}

func invert(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		b := s[i]
		var r byte
		switch b {
		case '0':
			r = 'f'
		case '1':
			r = 'e'
		case '2':
			r = 'd'
		case '3':
			r = 'c'
		case '4':
			r = 'b'
		case '5':
			r = 'a'
		case '6':
			r = '9'
		case '7':
			r = '8'
		case '8':
			r = '7'
		case '9':
			r = '6'
		case 'a':
			r = '5'
		case 'b':
			r = '4'
		case 'c':
			r = '3'
		case 'd':
			r = '2'
		case 'e':
			r = '1'
		case 'f':
			r = '0'
		}
		result += string([]byte{r})
	}
	return result
}

func increment(s string) string {
	result := make([]byte, len(s))
	carryOver := true
	for i := 0; i < len(s); i++ {
		a := s[len(s)-i-1]
		b, c := inc(a, carryOver)
		carryOver = c
		result[len(s)-i-1] = b
	}
	return string(result)
}

func inc(b byte, c bool) (byte, bool) {
	if !c {
		return b, false
	}
	if b == 'f' {
		return '0', true
	}
	if b == '9' {
		return 'a', false
	}
	return b + 1, false
}

func toHexdecimal(num int) string {
	if num < 10 {
		return fmt.Sprint(num)
	}
	return string([]byte{byte(num-10) + 97})
}
