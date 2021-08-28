package main

func intToRoman(num int) string {

	foursMap := map[int]string{
		4: "IV",
		9: "IX",
		40: "XL",
		90: "XC",
		400: "CD",
		900: "CM",
	}
	item := map[int]string{
		1: "I",
		5: "V",
		10: "X",
		50: "L",
		100: "C",
		500: "D",
		1000: "M",
	}

	standards := []int{1000, 500, 100, 50, 10, 5, 1}
	fours := []int{900, 400, 90, 40, 9, 4}

	result := ""
	
	current := num

	length := len(standards)

	for i := 0; i < length; i++ {
		standard := standards[i]
		x := current / standard
		for j := 1; j <= x; j++ {
			result = result + item[standard]
		}
		current = current - standard * x
		if standard != 1 && current >= fours[i] {
			current = current - fours[i]
			result = result + foursMap[fours[i]]
		}
	}

	return result
}
