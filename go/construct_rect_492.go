package main

import (
	"math"
)

/* solved on Sat Jul 1 2023 */
/*
	https://leetcode.com/problems/construct-the-rectangle/submissions/984113396/
*/

func constructRectangle(area int) []int {
	val := make([]int, 2)
	val[0] = area
	val[1] = 1

	sqrt := math.Sqrt(float64(area))

	for i := 1; i <= int(sqrt); i++ {
		if area%i == 0 {
			a := area / i
			if a >= i && a-i <= val[0]-val[1] {
				val[0] = a
				val[1] = i
			}
		}
	}

	return val
}
