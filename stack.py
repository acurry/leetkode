class Stack:
	def __init__(self):
		self.stack = []
	def push(self, s):
		self.stack.append(s)
	def pop(self):
		return self.stack.pop()
	def last(self):
		return self.stack[-1]
