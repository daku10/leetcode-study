package main

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	var tokens []string
	var str strings.Builder
	muldivCount := 0
	for _, r := range s {
		switch r {
		case ' ':
			continue
		case '+', '-', '*', '/':
			tokens = append(tokens, str.String())
			str.Reset()
			tokens = append(tokens, string(r))
			if r == '*' || r == '/' {
				muldivCount++
			}
		default:
			str.WriteByte(byte(r))
		}
	}
	tokens = append(tokens, str.String())
	for muldivCount != 0 {
	L1:
		for i, t := range tokens {
			switch t {
			case "*", "/":
				pre := toInt(tokens[i-1])
				nex := toInt(tokens[i+1])
				if t == "*" {
					tokens[i-1] = strconv.Itoa(pre * nex)
				} else {
					tokens[i-1] = strconv.Itoa(pre / nex)
				}
				copy(tokens[i:], tokens[i+2:])
				tokens = tokens[:len(tokens)-2]
				muldivCount--
				break L1
			}
		}
	}
	for len(tokens) != 1 {
	L2:
		for i, t := range tokens {
			switch t {
			case "+", "-":
				pre := toInt(tokens[i-1])
				nex := toInt(tokens[i+1])
				if t == "+" {
					tokens[i-1] = strconv.Itoa(pre + nex)
				} else {
					tokens[i-1] = strconv.Itoa(pre - nex)
				}
				copy(tokens[i:], tokens[i+2:])
				tokens = tokens[:len(tokens)-2]
				break L2
			}
		}
	}
	return toInt(tokens[0])
}

func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
