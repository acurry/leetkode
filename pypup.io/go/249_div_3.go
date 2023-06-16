package main

func div_3(x int) bool {
	if x == 0 || x == 3 || x == 6 || x == 9 {
		return true
	}

	a := 0
	val := false
	for a <= x {
		a += 3
		if a == x {
			val = true
			break
		}
	}

	return val
}
