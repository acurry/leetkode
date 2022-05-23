class ListNode:
	def __init__(self, val=0, next=None):
		self.val = val
		self.next = next

def addTwoNumbers(l1, l2):
	a_stack = []
	b_stack = []
	def iterate(node, stack):
		curr = node
		next = curr.next
		stack.append(curr.val)
		while next is not None:
			curr = next
			stack.append(curr.val)
			next = curr.next
			
	iterate(l1, a_stack)
	iterate(l2, b_stack)
	
	a = int(''.join(reversed([str(x) for x in a_stack])))
	b = int(''.join(reversed([str(x) for x in b_stack])))
	
	return a + b

a1 = ListNode(2, ListNode(4, ListNode(3)))
a2 = ListNode(5, ListNode(6, ListNode(4)))

