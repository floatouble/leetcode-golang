func lengthOfLastWord(s string) int {
    lastWordLength, currentWordLength := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			currentWordLength = 0
		} else {
			currentWordLength++
			lastWordLength = currentWordLength
		}
	}

	return lastWordLength
}