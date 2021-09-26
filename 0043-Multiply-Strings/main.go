package main

import (
	"fmt"
	"strings"
)

type MyDecimal struct {
    current []byte
    tenRate int
}

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    len1 := len(num1)
    len2 := len(num2)
    decimals := []MyDecimal{}
    for i := 0; i < len1; i++ {
        var over byte = 0
        // decimal has result of num2 * num1[x] in reverse order
        // ex. num2: 123, num1[x]: 2 => 6, 4, 2
        // the reason why order is reversed is convinient to append number.
        decimal := MyDecimal{[]byte{}, i}
        for j := 0; j < len2; j++ {
            // 48 means byte to int
            n1 := num1[len1 - i - 1] - 48
            n2 := num2[len2 - j - 1] - 48
            r := n1 * n2 + over
            one := r % 10
            decimal.current = append(decimal.current, one)
            over = r / 10
        }
        if over > 0 {
            decimal.current = append(decimal.current, over)
        }
        decimals = append(decimals, decimal)
    }
    
    return calculate(decimals)
}

func calculate(decimals []MyDecimal) string {

    dLen := len(decimals)

    arr := make([]byte, 0)

    for i := 0; i < dLen; i++ {
        decimal := decimals[i]
        currentLen := len(decimal.current)
        arrLen := len(arr)
        tmp := make([]byte, 0)

        for j := 0; j < decimal.tenRate; j++ {
            tmp = append(tmp, arr[j])
        }
        var over byte = 0
        for j := 0; j < currentLen; j++ {
            t := j + decimal.tenRate
            var anum byte = 0
            if t < arrLen {
                anum = arr[t]
            }
            dnum := decimal.current[j]
            r := anum + dnum + over
            one := r % 10
            tmp = append(tmp, one)
            over = r / 10
        }
        if over > 0 {
            tmp = append(tmp, over)
        }
        arr = tmp
    }

    arrLen := len(arr)

    result := strings.Builder{}
    for i := 0; i < arrLen; i++ {
        result.WriteString(fmt.Sprint(arr[arrLen - i - 1]))
    }

    return result.String()
}
