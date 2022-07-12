func trap(height []int) int {
    if  len(height) == 0 {
		return 0
	}
	
	left := 0
	right := len(height)-1
	leftMax := height[left]
	rightMax := height[right]
	res := 0
	
	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			res += (leftMax-height[left])
		} else {
			right--
			rightMax = max(rightMax, height[right])
			res += (rightMax-height[right])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} 
	return b
}