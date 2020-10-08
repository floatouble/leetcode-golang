import "unicode"

func isPalindrome(s string) bool {
    toLower := unicode.ToLower
    isLetterOrDigit := func (c rune) bool {
        return unicode.IsNumber(c) || unicode.IsLetter(c)
    }
    
    i := 0
    j := len(s)-1
    
    for i <= j {
        for i <=j && !isLetterOrDigit(rune(s[i])) {
            i++
        }
        
        for i <= j && !isLetterOrDigit(rune(s[j])) {
            j--
        }
        
        if i <= j && toLower(rune(s[i])) != toLower(rune(s[j])) {
            return false
        }
        i++
        j--
    }
    return true
    
}