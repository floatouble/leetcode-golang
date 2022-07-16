func checkInclusion(s1 string, s2 string) bool {
    left := 0
	right := 0

	freq := [26]int{}
	for _, r := range s1 {
		freq[r-'a']++
	}

	count := [26]int{}
	for right < len(s2) {
		count[s2[right]-'a']++
		if right-left+1 == len(s1) {
			if count == freq {
				return true
			}
			count[s2[left]-'a']--
			left++
		}
		right++
	}
	return false
}