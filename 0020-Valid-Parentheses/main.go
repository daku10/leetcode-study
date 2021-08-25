package main

func isValid(s string) bool {

	leftParenthesis := map[byte]struct{}{
		'(': {},
		'[': {},
		'{': {},
	}

	sLen := len(s)
	queue := make([]byte, 0)

	for i := 0; i < sLen; i++ {
		current := s[i]

		if _, ok := leftParenthesis[current]; ok {
			queue = append(queue, current)
		} else {
			qLen := len(queue)
			if qLen == 0 {
				return false
			}
			last := queue[qLen-1]
			if (current == ']' && last == '[') || (current == ')' && last == '(') || (current == '}' && last == '{') {
				queue = queue[0:qLen-1]
			} else {
				return false
			}
		}
	}

	return len(queue) == 0
}
