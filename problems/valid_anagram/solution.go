func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	alphas := make([]int, 26)
	for _, a := range s {
		alphas[a-'a']++
	}

	for _, a := range t {
		if alphas[a-'a'] == 0 {
			return false
		}

		alphas[a-'a']--
	}

	for _, a := range alphas {
		if a != 0 {
			return false
		}
	}

	return true
}