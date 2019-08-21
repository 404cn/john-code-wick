#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#
class Solution:
    """
    # 3004ms, 22.98%, 14.41%(14.8MB)
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            x = target - nums[i]
            for j in range(i + 1, len(nums)):
                if x == nums[j]:
                    return i, j

    """
    # 56ms, 87.03%, 5.11%(15.2MB)
    def twoSum(self, num, target):
        map = {}
        for i in range(len(num)):
            if num[i] not in map:
                map[target - num[i]] = i
            else:
                return map[num[i]], i
        return -1, -1


