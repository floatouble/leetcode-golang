func isPalindrome(s string) bool {
    i := 0
	j := len(s) - 1
    
    letterOrDigit:=func(x rune) bool{
        return unicode.IsLetter(x) || unicode.IsDigit(x)
    }

	s = strings.ToLower(s)
	for i <= j {
		for i <= j && !letterOrDigit(rune(s[i])) {
			i++
		}

		for i <= j && !letterOrDigit(rune(s[j])) {
			j--
		}

		if i <= j && s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}