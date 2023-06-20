package main

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	a, b, c := 0, 1, 1

	for i := 0; i < n-1; i++ {
		a = b
		b = c
		c = a + b
	}

	return c

	// return climbStairs(n-1) + climbStairs(n-2)

	/*
		0 1 1 2 3 5 8 13

		4 = n + 1
		1 1 1 1
		1 1 2
		1 2 1
		2 1 1
		2 2

		5 = n + 3
		1 1 1 1 1
		1 1 1 2
		1 1 2 1
		1 2 1 1
		2 1 1 1
		1 2 2
		2 1 2
		2 2 1

		6
		1 1 1 1 1 1
		1 1 1 1 2
		1 1 1 2 1
		1 1 2 1 1
		1 2 1 1 1
		2 1 1 1 1
		1 1 2 2
		1 2 2 1
		2 2 1 1
		2 2 2
	*/
}
