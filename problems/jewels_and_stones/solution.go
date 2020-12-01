func numJewelsInStones(J string, S string) int {
    jewels := make(map[rune]struct{}, len(J))
    
    for _,c := range J {
        jewels[c] = struct{}{}
    }
    
    jewelCount := 0
    for _,c := range S {
        if _, ok := jewels[c]; ok {
            jewelCount++
        }
    }
    return jewelCount
    
}