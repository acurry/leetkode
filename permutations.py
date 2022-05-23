def list_swap(l, i, j):
	a, b = l[i], l[j]
	l[i] = b
	l[j] = a
	

def permute(nums):
	ret_val = []
	
	def gen(k, l, ret_val):
		if k == 1:
			ret_val.append(l)
			return ret_val
		else:
			for i in range(k):
				gen(k-1, l)
				if k % 2 == 0:
					list_swap(l, i, k-1)
				else:
					list_swap(l, 0, k-1)
					
	gen(len(nums), nums)
	return ret_val

a = [1,2,3]
b = [0,1]
c = [1]
