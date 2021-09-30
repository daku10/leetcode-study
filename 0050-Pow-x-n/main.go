package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MyDecimal struct {
	arr []int
	scale int
	sign int // 1 or -1
}

func product(l MyDecimal, r MyDecimal) MyDecimal {
	newScale := l.scale + r.scale
	newSign := l.sign * r.sign
	lArr := l.arr
	rArr := r.arr
	lLength := len(lArr)
	rLength := len(rArr)

	
	result := make([]int, 0)
	for i := 0; i < lLength; i++ {
		arr := make([]int, 0)
		over := 0
		for j := 0; j < rLength; j++ {
			current := lArr[i] * rArr[j] + over
			over = current / 10
			current = current % 10
			arr = append(arr, current)
		}
		if over > 0 {
			arr = append(arr, over)
		}
		lenArr := len(arr)

		newResult := make([]int, 0)
		lenResult := len(result)
		for j := 0; j < i; j++ {
			newResult = append(newResult, result[j])
		}
		over = 0		
		for j := 0; j < lenArr; j++ {
			target := 0			
			if j + i < lenResult {
				target = result[j + i]
			}
			current := target + arr[j] + over
			over = current / 10
			current = current % 10
			newResult = append(newResult, current)
		}
		if over > 0 {
			newResult = append(newResult, over)
		}
		result = newResult
	}

	return MyDecimal{result, newScale, newSign}
}

func createDecimal(f float64) MyDecimal {
	sign := 1
	if f < 0 {
		f = -f
		sign = -1
	}
	s := strconv.FormatFloat(f, 'f', -1, 64)
	index := strings.Index(s, ".")
    scale := 0
	if index == -1 {
        scale = 0
    } else {
		scale = len(s[index+1:])
	}
	s = strings.ReplaceAll(s, ".", "")
	arr := make([]int, 0)
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		// 48 is bias of byte to int
		arr = append(arr, int(s[i] - 48))
	}
	return MyDecimal{arr, scale, sign}
}

func (d MyDecimal) toString() string {

	builder := strings.Builder{}

	arrLen := len(d.arr)

	if d.sign == -1 {
		builder.WriteString("-")
	}

	for i := 0; i < arrLen; i++ {
		if i == (arrLen - d.scale) {
			builder.WriteString(".")
		}
		builder.WriteString(fmt.Sprint(d.arr[arrLen - i - 1]))
	}
	return builder.String()
}

func myPow(x float64, n int) float64 {

	d := createDecimal(x)

	isMinus := false
	if n < 0 {
		n = -n
		isMinus = true
	}

	result := createDecimal(1)
	for i := 0; i < n; i++ {
		result = product(result, d)
	}

	s := result.toString()
	f, _ := strconv.ParseFloat(s, 64)
	if isMinus {
		return 1 / f
	}
	return f
}
