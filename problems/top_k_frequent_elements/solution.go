func topKFrequent(nums []int, k int) []int {
	counters := make(map[int]int)
	frequency := make([][]int, len(nums)+1)
	for _, n := range nums {
		counters[n]++
	}

	for k, v := range counters {
		frequency[v] = append(frequency[v], k)
	}

	res := []int{}
	for i := len(frequency) - 1; i >= 0; i-- {
		for _, n := range frequency[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}