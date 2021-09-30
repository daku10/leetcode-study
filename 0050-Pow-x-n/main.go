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

	if n == 0 {
		return 1.0
	}

	d := createDecimal(x)

	isMinus := false
	needOneMore := false
	if n < 0 {
		// so stupid logic
		if n == -2147483648 {
			n = 2147483647
			isMinus = true
			needOneMore = true
		} else {
			n = -n
			isMinus = true
		}
	}

	result := createDecimal(1)

	current := n
	rest := 0
	limitScale := 20

	for {
		m, l := destructPow(current)
		md := createDecimal(x)
		for i := 0; i < m; i++ {
			md = product(md, md)
			md.floorScale(limitScale)
			// so badly extremely ugly...
			// the problem declares -10^4 <= xn <= 10^4
			// if md.arr grows larger, it should be (1 / large number), should be zero.
			// this logic is so bad.
			if md.scale == 0 && len(md.arr) > 100 {
				return 0.0
			}
		}
		result = product(result, md)
		result.floorScale(limitScale)
		current = l
		if l < 10 {
			rest = l
			break;
		}
	}

	for i := 0; i < rest; i++ {
		result = product(result, d)
	}

	if needOneMore {
		result = product(result, d)
	}

	s := result.toString()
	f, _ := strconv.ParseFloat(s, 64)
	if isMinus {
		return 1 / f
	}
	return f
}

// n => 2^m + l
func destructPow(n int) (int, int) {
	if n == 0 {
		return 0, 0
	}
	m := 0
	current := 1
	for {
		if current * 2 > n {
			break
		}
		m++
		current = current * 2
	}
	l := n - current
	return m, l
}

func (d *MyDecimal) floorScale(scale int) {
	length := len(d.arr)
	if length < scale || d.scale < scale {
		return
	}
	d.arr = d.arr[(d.scale - scale):length]
	d.scale = scale
}
