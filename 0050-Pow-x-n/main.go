package main

import (
	"strconv"
	"strings"
)

func myPow(x float64, n int) float64 {

	d := FloatSignificantDigit(x)

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

func FloatSignificantDigit(f float64) int {
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
