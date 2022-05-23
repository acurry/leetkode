def longestCommonPrefix(strs):
	prefix = ""
	hashmap = {}
	for s in strs:
		
	
	if len(prefix) == 0:
		return "There is no common prefix among the input strings."
	else:
		return prefix

tests = [
	["flower", "flight", "flow"],
	["", "foo"],
	["dog", "cat", "racecar"],
]

def main():
	for test in tests:
		print(f"{test}: {longestCommonPrefix(test)}")

