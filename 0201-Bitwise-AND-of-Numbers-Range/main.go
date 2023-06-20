package main

func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left
	}

	if right-left == 1 {
		return right & left
	}

	max := 1073741824 // 2^30
	result := 0
	for i := 31; i > 0; i-- {
		l := left & (max >> (31 - i))
		r := right & (max >> (31 - i))
		if l != 0 || r != 0 {
			if l != r {
				return result
			}
			result += l
		}
	}

	return result
}
