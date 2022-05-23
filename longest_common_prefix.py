def longestCommonPrefix(strs):
	prefix = ""	
	shortest_length = sorted([len(x) for x in _strs])[0]
	index = 0
	for index in range(shortest_length):
		sl = set([x[index:index+1] for x in _strs])
		if len(sl) == 1:
			prefix += _strs[0][index]
		else:
			break
	
	if prefix == "":
		return "There is no common prefix among the input strings."

	return prefix


_strs = ["flower","flow","flight"]

