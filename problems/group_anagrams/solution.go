func groupAnagrams(input []string) [][]string {
    groups := make([][]string, 0)
	addedWords := make(map[string]struct{})

	for _, s := range input {
		if _, ok := addedWords[s]; ok {
			continue
		}

		g := make([]string, 0)
		for _, w := range input {
			if isAnagram2(s, w) {
				g = append(g, w)
				addedWords[w] = struct{}{}
			}
		}
		groups = append(groups, g)
	}

	return groups
    
}

func isAnagram2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	alphas := make([]int, 255)
	for _, a := range s1 {
		alphas[a]++
	}

	for _, a := range s2 {
		if alphas[a] == 0 {
			return false
		}

		alphas[a]--
	}

	for _, a := range alphas {
		if a != 0 {
			return false
		}
	}

	return true
}
