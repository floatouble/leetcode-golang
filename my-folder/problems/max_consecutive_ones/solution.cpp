class Solution {
public:
    int findMaxConsecutiveOnes(vector<int>& nums) {
        int count = 0;
    int maxCount = 0;

    for (int i = 0; i < nums.size(); i++)
    {
        if (nums[i] != 1)
        {
            maxCount = std::max(maxCount, count);
            count = 0;
        }
        else
        {
            count++;
        }
    }
    return std::max(maxCount, count);
    }
};