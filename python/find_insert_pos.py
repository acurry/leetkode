class Solution:
	def bs(self, nums, target):
		first = 0
		last = len(nums) - 1
		mid = (first+last) // 2
		while first <= last:
			mid = (first+last) // 2
			if target == nums[mid]:
				return mid
			elif target < nums[mid]:
				last = mid - 1
			elif target > nums[mid]:
				first = mid + 1
			if nums[mid] <= target and nums[mid] > target:
				return mid
			else:
				return mid+1

	def search(self, nums, target):
		if target < nums[0]:
			return 0
		elif target > nums[-1]:
			return len(nums)
		
		return self.bs(nums, target)
			
a = [1,3]
b = 2

s = Solution()

