package main

// https://pypup.com/problems/frequent-char

func frequent_char(s string) string {
	counts := make(map[string]int, len(s))
	biggest := ""

	for _, c := range s {
		cc := string(c)
		counts[cc] += 1
		if counts[cc] > counts[biggest] {
			biggest = cc
		}
	}

	return string(biggest)
}
