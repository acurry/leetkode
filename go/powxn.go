package main

import (
	"fmt"
	"math"
)

func MyPow(x float64, n int) float64 {
	val := x
	if x == 1.0 {
		return val
	} else if x == -1.0 {
		if n%2 != 0 {
			return val
		} else {
			return 0 - val
		}
	}

	if n == 0 {
		if x > 0 {
			return 1.0
		} else {
			return val
		}
	}

	bound := int(math.Abs(float64(n)))

	for i := 1; i < bound; i++ {
		val = val * x
	}

	if n < 0 {
		return 1 / val
	} else {
		return val
	}
}

func main() {
	fmt.Println("MyPow(-1, 2): ", MyPow(-1, 2))
	fmt.Println("MyPow(-1, 5): ", MyPow(-1, 5))
	fmt.Print("SOLVED!")
}
