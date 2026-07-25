func topKFrequent(nums []int, k int) []int {
	 // key = number, value = frequency count
	hashTable := make(map[int]int)
	for _, v := range nums {
		hashTable[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range hashTable {
		buckets[count] = append(buckets[count], num)
	}

	res := []int{}
	for i := len(nums); i>0; i-- {
		if len(buckets[i]) == 0 {
			continue
		}

		res = append(res, buckets[i]...)
		
		if len(res) >= k {
			break
		}	
	}

	return res[:k]
}
