package main

func plusOne(digits []int) []int {
	x := prependZero(digits)

	length := len(x)

	for i := length - 1; i >= 0; i-- {
		a := x[i] + 1

		if a < 10 {
			x[i] = a
			break
		} else {
			x[i] = 0
		}
	}

	if x[0] == 0 {
		x = x[1:]
	}

	return x
}

func prependZero(a []int) []int {
	x := make([]int, len(a))
	// sum = [1, 2, 3]
	copy(x, a)

	// sum = [3,2,1]
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	// x = [3,2,1,0]
	x = append(x, 0)

	// x = [0, 1, 2, 3]
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	return x
}
