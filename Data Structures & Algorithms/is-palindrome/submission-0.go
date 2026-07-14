func isPalindrome(s string) bool {
	runes := []rune(strings.ToLower(s))
	i := 0
	j := len(runes)-1
	for i<j {
		if !isAlphaNumeric(runes[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(runes[j]) {
			j--
			continue
		}

		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || 
		(r >= 'A' && r <= 'Z') || 
		(r >= '0' && r <= '9')
}
