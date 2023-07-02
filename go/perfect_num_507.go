package main

// https://leetcode.com/problems/perfect-number/submissions/984116586/

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	midpoint := num / 2
	sum := 1

	for i := 2; i <= midpoint; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num
}
