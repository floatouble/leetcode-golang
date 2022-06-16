func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))
    for i, _ := range result {
        result[i] = 1
    }
    prefix := 1
    for i, v := range nums {
        result[i] *= prefix
        prefix *= v
    }
    
    postfix := 1
    for i := len(nums)-1; i>=0; i-- {
        result[i] *= postfix
        postfix *= nums[i]
    }
    
    return result
}