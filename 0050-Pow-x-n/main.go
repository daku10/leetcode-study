package main

import (
	"strconv"
	"strings"
)

type MyDecimal struct {
	arr []int
	scale int
}

func product(l MyDecimal, r MyDecimal) MyDecimal {
	newScale := l.scale + r.scale
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

	return MyDecimal{result, newScale}
}

func createDecimal(f float64) MyDecimal {
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
	return MyDecimal{arr, scale}
}

func myPow(x float64, n int) float64 {

	d := floatSignificantDigit(x)

	var geta int64 = 1
	for i := 0; i < d; i++ {
		geta *= 10
	}

	getaedX := attachGeta(x, d)

	var result int64 = 1
	var bigGeta int64 = 1

	isMinus := false

	if n < 0 {
		n = -n
		isMinus = true
	}

	for i := 0; i < n; i++ {
		result *= getaedX
		bigGeta *= geta
	}


	if isMinus {
		return 1 / (float64(result) / float64(bigGeta))
	}

	return float64(result) / float64(bigGeta)
}

func floatSignificantDigit(f float64) int {
    s := strconv.FormatFloat(f, 'f', -1, 64)

    i := strings.Index(s, ".")
    if i == -1 {
        return 0
    }

    return len(s[i+1:])
}

func attachGeta(x float64, d int) int64 {
	s := strconv.FormatFloat(x, 'f', -1, 64)
	s = strings.Replace(s, ".", "", -1)
    n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
