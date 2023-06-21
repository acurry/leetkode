class Solution:
    def myPow(self, x: float, n: int) -> float:
        val = x

        if n == 0:
            if x > 0:
                return 1.0
            else:
                return val

        for _ in range(1, abs(n)):
            val *= x
        if n < 0:
            return 1.0 / val
        else:
            return val


### latest run on leetcode:

# Time Limit Exceeded
# 291 / 305 testcases passed
# Last Executed Input
# Use Testcase
# x =
# 0.00001
# n =
# 2147483647
