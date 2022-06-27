func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for _, v := range nums {
        set[v] = struct{}{}
    }
    
    maxLength := 0
    for _, v := range nums {
        if _, ok := set[v-1]; !ok {
            length := 1
            i := 1
            for {
                if _, ok := set[v+i]; !ok {
                    break
                }
                i++
                length++
            }
            maxLength = max(maxLength, length)
        }
    }
    return maxLength
}


func max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}