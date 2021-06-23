class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        arr = {}
        for index, value in enumerate(nums):
            diff = target - value
            if diff in arr:
                return [arr[diff], index]
            else:
                arr[value] = index
        return ('Cannot find the valid indices')