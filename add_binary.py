def to_i(b):
	res = 0
	length = len(b)
	for i in range(length):
		if b[i] == '1':
			res += 2 ** (length-i-1)
	return res

