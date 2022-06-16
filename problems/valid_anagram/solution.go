func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    counter := make(map[rune]int)
    for _, c := range s {
        counter[c]++
    }
    
    for _, c := range t {
        if _, ok := counter[c]; !ok {
            return false
        }
        counter[c]--
    }
    
    for _, v := range counter {
        if v != 0 {
            return false
        }
    }
    
    return true
}