#
# @lc app=leetcode id=7 lang=python3
#
# [7] Reverse Integer
#
class Solution:
    def reverse(self, x: int) -> int:
        y = 0
        if x > 2**32 - 1 or x < -2**32:
            return 0
#            y = y + () * 10**(i - 1)

