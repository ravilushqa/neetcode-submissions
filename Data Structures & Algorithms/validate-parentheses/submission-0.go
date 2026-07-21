func isValid(s string) bool {
	stack := []rune{}

	for _, v := range []rune(s) {
		switch v {
		case '(', '{','[':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if (v == ')' && top == '(') || (v == '}' && top == '{') || (v == ']' && top == '[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
