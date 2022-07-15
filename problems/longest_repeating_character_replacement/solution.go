func characterReplacement(s string, k int) int {
    left := 0
	right := 0
	
	maxLength := 0
	count := map[byte]int{}
	maxF := 0
	
	// maxLength = window_length - max_frequency_of_char <= k
	
	for right < len(s) {
		count[s[right]]++
		maxF = max(maxF, count[s[right]])
		
        for (right-left+1)-maxF > k {
			count[s[left]]--
			left++
		}
		maxLength = max(maxLength, right-left+1)
        right++
	}
	
	return maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b   
}