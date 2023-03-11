package main

func singleNumber(nums []int) int {
	a := 0
	b := 0
	for _, c := range nums {
		ta := (a & ^b & ^c) | (^a & b & c)
		b = (^a & b &^ c) | (^a &^ b & c)
		a = ta
	}
	return b
}
