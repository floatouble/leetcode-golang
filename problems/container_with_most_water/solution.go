func maxArea(height []int) int {
	var maxA int = math.MinInt
	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxA = max(area, maxA)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxA
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}