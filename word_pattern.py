def wordPattern(pattern, s):
	chars_to_words = {}
	chars = list(pattern)
	words = list(s.split(' '))
	val = True
	for (i, j) in zip(chars, words):
		if i not in chars_to_words:
			chars_to_words[i] = j
		if chars_to_words[i] != j:
			val = False
			break
	return val
		

a = 'aba'
b = 'dog cat cat'

