func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}
	
	left := 0
	right := 0
	
	minLength := math.MaxInt
	minString := ""
	
	window := map[byte]int{}

	countT := map[byte]int{}
	for i:= 0; i < len(t); i++ {
		countT[t[i]]++
	}
	
	// number of unique characters in a string
	// e.g. if t is "AABC" then the window must have two A's, one B
	// and one C. Thus formed would be = 3 when all these conditions
	// are met.
	needsCount := len(countT)
	havesCount := 0
	
	for right < len(s) {
		window[s[right]]++
		if _, ok := countT[s[right]]; ok {
			if window[s[right]] == countT[s[right]] {
				havesCount++
			}
		}
		
		for havesCount == needsCount { // for
			if minLength > (right-left+1) {
				minLength =  right-left+1
				minString = s[left : right+1]
			}
			
			// shrink the window
			window[s[left]]--
			if _, ok := countT[s[left]]; ok {
				if window[s[left]] < countT[s[left]] {
					havesCount--
				}
			}
			left++
		}
		right++
	}
	
	return func(minLength int, minString string) string {
		if minLength != math.MaxInt {
			return minString
		}
		return ""
	}(minLength, minString)
}

func min(a, b int)int {
	if a < b {
		return a
	}
	return b
}
