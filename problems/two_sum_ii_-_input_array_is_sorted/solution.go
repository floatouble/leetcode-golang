func twoSum(nums []int, target int) []int {
    i := 0
	j := len(nums) - 1

	for i < j {
		sum := nums[i] + nums[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{0, 0}
}