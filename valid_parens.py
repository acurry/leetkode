def isValid(s):
	if len(s) % 2 != 0:
		return False
	if s[0] in [')', '}', ']']:
		return False
	stack = []
	chars = list(s)
	index = 0
	while True:
		if s[index] in ['(', '{', '[']:
			stack.append(s[index])
		else:
			if len(stack) == 0:
				break
			last = stack[len(stack)-1]
			if f"{last}{s[index]}" in ['()', '{}', '[]']:
				stack.pop()
			else:
				break
		index += 1
	return len(stack) == 0 and index == len(s) - 1

tests = [
	'()',
	'({[]})',
	'}',
	'(){}}{'
]

def main():
	for test in tests:
		print(f"{test}: {isValid(test)}")

