def isIsomorphic(s, t):
	hashmap = {}
	for (i, j) in zip(s, t):
		if i not in hashmap:
			hashmap[i] = j
		if hashmap[i] != j:
			return False
	return True

a = 'egg'
b = 'add'
c = 'paper'
d = 'title'
e = 'foo'
f = 'bar'

