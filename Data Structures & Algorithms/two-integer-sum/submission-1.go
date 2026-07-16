func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, exists := m[target-v]; exists {
			return []int{j, i}
		}

		m[v] = i
	}

	return []int{}

}
