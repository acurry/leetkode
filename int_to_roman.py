def digit_to_roman(n, a, b, c):
	v = ''
	if n <= 3:
		v = n * a
	elif n == 4:
		v = f"{a}{b}"
	elif n >= 5 and n <= 8:
		v = b
		if n > 5 and n <= 8:
			v += (n - 5) * a
	elif n == 9:
		v = f"{a}{c}"
	return v

def ones(n):
	return digit_to_roman(n, 'I', 'V', 'X')
	
def tens(n):
	return digit_to_roman(n, 'X', 'L', 'C')

def hundreds(n):
	return digit_to_roman(n, 'C', 'D', 'M')
	
def thousands(n):
	return n * 'M'

def intToRoman(num): # -> str
	digits = list(reversed([int(x) for x in list(str(num))]))
	val = ""
	
	funcs = [ones, tens, hundreds, thousands]
	
	for i in range(len(digits)):
		num = funcs[i](digits[i])
		val = f"{num}{val}"

	return val

a = 3
b = 58
c = 1994

