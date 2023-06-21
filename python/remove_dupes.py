def removeDuplicates(nums):
	hashmap = {}
	for n in nums:
		if not n in hashmap:
			hashmap[n] = []
		hashmap[n].append(n)
	keys = list(hashmap.keys())
	diff = len(nums) - len(keys)
	for i in range(len(keys)):
		nums[i] = keys[i]
	for d in range(diff+1, len(nums)):
		nums[d] = '_'
	return len(keys)
		
#        keys = list(hashmap.keys())
#        diff = len(nums) - len(keys)
#        for i in range(len(keys)):
#            nums[i] = keys[i]
#        for d in range(diff, len(nums)):
#            nums[d] = '_'           
#        return len(hashmap.keys())
