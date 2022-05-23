MAX = (2 ** 31) - 1

def reverse(self, x: int) -> int:
	val = 0
	x_chars = list(str(x))
	
	if x_chars[0] == '-':
		reversed_x = int(''.join(reversed(x_chars[1:])))
		if reversed_x <= MAX:
			val = int(''.join(list('-') + reversed_x))
	else:
		reversed_x = int(''.join(reversed(x_chars)))
		if reversed_x <= MAX:
			val = int(''.join(reversed_x))
		
	return val
			
		

