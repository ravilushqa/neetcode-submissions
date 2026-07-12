func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
        return false
    }

	firstMap  := make(map[rune]int, 26)
	secondMap := make(map[rune]int, 26)

	for _, v := range s{
		firstMap[v]++
	}
	for _, v := range t{
		secondMap[v]++
	}
	
	if len(firstMap) != len(secondMap) {
		return false
	}

	for k, v := range firstMap {
		if secondMap[k] != v {
			return false
		}
	}

	return true
}
