def lengthOfLongestSubstring(s):
	if len(s) == 0:
		return 0
		
	substrs = [[]]
	substr_count = 0
	for i in range(len(s)):
		if s[i] not in substrs[substr_count]:
			substrs[substr_count].append(s[i])
		else:
			substr_count += 1
			substrs.append([s[i]])
	
	return sorted([len(x) for x in substrs])[-1]

a = 'abcabcbb'

