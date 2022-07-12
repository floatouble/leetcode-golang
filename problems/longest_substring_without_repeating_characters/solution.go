func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
		return 0
	}

	left := 0
	right := 0
	maxLength := 0

	m := map[byte]struct{}{}
	for left < len(s) && right < len(s) {
		if _, ok := m[s[right]]; ok {
			maxLength = max(maxLength, right-left)
			for {
				if _, ok := m[s[right]]; ok {
					delete(m, s[left])
					left++
				} else {
					break
				}
			}
		} else {
			m[s[right]] = struct{}{}
			right++
		}
	}
	return max(maxLength, right-left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}