num_to_alpha = {
	'2': ['a', 'b', 'c'],
	'3': ['d', 'e', 'f'],
	'4': ['g', 'h', 'i'],
	'5': ['j', 'k', 'l'],
	'6': ['m', 'n', 'o'],
	'7': ['p', 'q', 'r', 's'],
	'8': ['t', 'u', 'v'],
	'9': ['w', 'x', 'y', 'z'],
}
def letterCombinations(digits):
	if len(digits) == 0:
		return []
	result = []
	
	letters = []
	for i in digits:
		letters += num_to_alpha[i]

	return letters

