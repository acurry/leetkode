package main

import "fmt"

/*
Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
*/

func fizzBuzz(n int) []string {
	vals := make([]string, n)

	for i := 0; i < n; i++ {
		a := i + 1
		if a%5 == 0 && a%3 == 0 {
			vals[i] = "FizzBuzz"
		} else if a%5 == 0 {
			vals[i] = "Buzz"
		} else if a%3 == 0 {
			vals[i] = "Fizz"
		} else {
			vals[i] = fmt.Sprint(a)
		}
	}

	return vals
}
