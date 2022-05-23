class Solution:
	def twoSum(self, nums, target):
		hashmap = {}
		result = []
		for i in range(len(nums)):
			diff = target - nums[i]
			if diff not in hashmap:
				hashmap[diff] = i
			else:
				result = [hashmap[diff], i]
				break


		return result


s = Solution()

a = [2,7,11,15]
b = 9

