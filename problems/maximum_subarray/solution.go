func maxSubArray(nums []int) int {
    maxHere := 0
    maxSoFar := math.MinInt
    
    for _, v := range nums {
        maxHere = max(v, v+maxHere)
        maxSoFar = max(maxSoFar, maxHere)
    }
    
    return maxSoFar
}

func max(x, y int) int {
    if x >= y {
        return x
    }
    
    return y
}