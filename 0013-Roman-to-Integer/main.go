package main


func romanToInt(s string) int {

	romanIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	decreaseMap := map[byte]map[byte]struct{}{
		'I': {'V': {}, 'X': {}},
		'X': {'L': {}, 'C': {}},
		'C': {'D': {}, 'M': {}},
	}

	result := 0

	sLen := len(s)
	for i := 0; i < sLen; i++ {
		current := s[i]
		if _, exists := decreaseMap[current]; exists && i + 1 != sLen {
			next := s[i + 1]
			if _, exists := decreaseMap[current][next]; exists {
				minus := romanIntMap[current]
				value := romanIntMap[next]
				result = result + value - minus
				i++
				continue;
			}
		}
		result = result + romanIntMap[current]
	}
	
	return result
}