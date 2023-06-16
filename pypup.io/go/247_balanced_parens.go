package main

func balanced_paranthesis(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	l := 0
	r := 0
	for _, c := range s {
		if c == '(' {
			l++
		} else {
			r++
		}
	}

	return l == r
}
