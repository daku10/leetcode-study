package main

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	var tokens []string
	var str strings.Builder
	for _, r := range s {
		switch r {
		case ' ':
			continue
		case '+', '-', '*', '/':
			tokens = append(tokens, str.String())
			str.Reset()
			tokens = append(tokens, string(r))
		default:
			str.WriteByte(byte(r))
		}
	}
	tokens = append(tokens, str.String())
	used := make(map[int]struct{})
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		switch t {
		case "*", "/":
			prevIndex := i - 1
			for {
				if _, ok := used[prevIndex]; !ok {
					break
				}
				prevIndex--
			}
			pre := toInt(tokens[prevIndex])
			nex := toInt(tokens[i+1])
			if t == "*" {
				tokens[i+1] = strconv.Itoa(pre * nex)
			} else {
				tokens[i+1] = strconv.Itoa(pre / nex)
			}
			used[i] = struct{}{}
			used[prevIndex] = struct{}{}
			i = i + 1
		}
	}
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		switch t {
		case "+", "-":
			prevIndex := i - 1
			for {
				if _, ok := used[prevIndex]; !ok {
					break
				}
				prevIndex--
			}
			nextIndex := i + 1
			for {
				if _, ok := used[nextIndex]; !ok {
					break
				}
				nextIndex++
			}
			pre := toInt(tokens[prevIndex])
			nex := toInt(tokens[nextIndex])
			if t == "+" {
				tokens[nextIndex] = strconv.Itoa(pre + nex)
			} else {
				tokens[nextIndex] = strconv.Itoa(pre - nex)
			}
			used[i] = struct{}{}
			used[prevIndex] = struct{}{}
			i = nextIndex
		}
	}

	index := 0
	for {
		if _, ok := used[index]; !ok {
			break
		}
		index++
	}
	return toInt(tokens[index])
}

func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
