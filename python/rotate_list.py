class ListNode:
	def __init__(self, val=0, next=None):
		self.val = val
		self.next = next
	def __str__(self):
		return f"Node_{self.val}"
		
def rotateRight(head, k):
	length = 0
	curr = head
	nex = None
	new_head = None
	
	# compute length of the list
	while curr != None:
		nex = curr
		curr = nex.next
		length += 1
	offset = k % length
	
	curr = head
	nex = None
	for i in range(offset):
		

e = ListNode(5)
d = ListNode(4, e)	
c = ListNode(3, d)
b = ListNode(2, c)
a = ListNode(1, b)

