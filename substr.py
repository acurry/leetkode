def strStr(haystack, needle):
	if len(needle) == 0 or len(haystack) == 0:
		return -1
	i = 0
	j = len(needle)
	index = -1
	while j <= len(haystack):
		if needle == haystack[i,j]:
			index = i
			break
		else
			i += 1
			j += 1
	return index

