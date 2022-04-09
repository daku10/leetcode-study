package main

func addBinary(a string, b string) string {
	carryOver := false
	lenA := len(a)
	lenB := len(b)
	maxLength := lenA
	if maxLength < lenB {
		maxLength = lenB
	}

	res := ""

	for i := 0; i < maxLength; i++ {
		ai := "0"
		if i < lenA {
			ai = string(a[lenA-i-1])
		}
		bi := "0"
		if i < lenB {
			bi = string(b[lenB-i-1])
		}
		r, c := add(ai, bi, carryOver)
		carryOver = c
		res = r + res
	}

	if carryOver {
		return "1" + res
	}

	return res
}

func add(a string, b string, carryOver bool) (string, bool) {
	if a == b {
		if carryOver {
			// 1, 1, true
			if a == "1" {
				return "1", true
				// 0, 0, true
			} else {
				return "1", false
			}
		} else {
			// 1, 1, false
			if a == "1" {
				return "0", true
				// 0, 0, false
			} else {
				return "0", false
			}
		}
	} else {
		// 1, 0, true
		if carryOver {
			return "0", true
			// 1, 0, false
		} else {
			return "1", false
		}
	}
}
