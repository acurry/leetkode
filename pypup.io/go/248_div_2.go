package main

import "fmt"

func div_2(x int) bool {
	s := fmt.Sprintf("%d", x)

	return s[len(s)-1] == '0' || s[len(s)-1] == '2' || s[len(s)-1] == '4' || s[len(s)-1] == '6' || s[len(s)-1] == '8'
}
