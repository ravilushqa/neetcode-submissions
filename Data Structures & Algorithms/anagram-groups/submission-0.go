func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)

	groupedHash := make(map[[26]int][]string)

	for _, str := range strs {
		var ar [26]int
		
		for _, char := range []rune(str) {
			ar[char - 'a']++
		}

		groupedHash[ar] = append(groupedHash[ar], str)
	}

	for _, v := range groupedHash {
		res = append(res, v)
	}

	return res
}
