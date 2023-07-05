package main

import "math"

func maxArea(height []int) int {
	max := 0
	length := len(height)

	for i := 0; i < length; i++ {
		currHeight := height[i]

		for j := i + 1; j < length; j++ {
			secondCurrHeight := height[j]
			currContainerLength := int(math.Abs(float64(i - j)))
			shorterContainer := secondCurrHeight
			if secondCurrHeight > currHeight {
				shorterContainer = currHeight
			}

			areaToCheck := currContainerLength * shorterContainer

			if areaToCheck > max {
				max = int(areaToCheck)
			}
		}
	}

	return max
}
