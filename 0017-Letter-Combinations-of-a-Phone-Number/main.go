package main

func letterCombinations(digits string) []string {

	itemMap := map[byte][]string {
		'2' : {"a", "b", "c"},
		'3' : {"d", "e", "f"},
		'4' : {"g", "h", "i"},
		'5' : {"j", "k", "l"},
		'6' : {"m", "n", "o"},
		'7' : {"p", "q", "r", "s"},
		'8' : {"t", "u", "v"},
		'9' : {"w", "x", "y", "z"},
	}

	length := len(digits)

	queue := []string{""}

	for i := 0; i < length; i++ {
		digit := digits[i]
		candidates := itemMap[digit]

		queueLength := len(queue)
		tmp := make([]string, 0, queueLength * 4)
		for j := 0; j < queueLength; j++ {
			q := queue[j]

			lenCandidate := len(candidates)
			for k := 0; k < lenCandidate; k++ {
				ap := candidates[k]
				tmp = append(tmp, q + ap)
			}
		}
		queue = tmp
	}

	if len(queue) == 1 {
		return []string{}
	}

	return queue
}
