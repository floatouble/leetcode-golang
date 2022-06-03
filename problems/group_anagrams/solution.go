func groupAnagrams(strs []string) [][]string {
    mm := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for _, v := range str {
			key[v-'a']++
		}

		mm[key] = append(mm[key], str)
	}

	arrs := [][]string{}
	for _, v := range mm {
		arrs = append(arrs, v)
	}
	return arrs
}
