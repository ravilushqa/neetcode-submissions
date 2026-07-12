func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
        return false
    }

	firstMap  := make(map[rune]int, 26)
	secondMap := make(map[rune]int, 26)

	for i, v := range s{
		firstMap[v]++
		secondMap[rune(t[i])]++
	}

	for k, v := range firstMap {
		if secondMap[k] != v {
			return false
		}
	}

	return true
}
