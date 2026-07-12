func isAnagram(s string, t string) bool {
	firstMap  := make(map[rune]int, 0)
	secondMap := make(map[rune]int, 0)

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
