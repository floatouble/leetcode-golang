func maxProfit(prices []int) int {
    if len(prices) <= 1 {
		return 0
	}
	
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	
	maxxProfit := 0
	left := 0
	right := left+1
	
	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right]-prices[left]
			maxxProfit = max(profit, maxxProfit)
		} else {
			left = right
		}
		right++
	}
	return maxxProfit
}