package main

func unique_chars_count(s string) int {
	counts := map[string]int{}
	for _, c := range s {
		cc := string(c)
		counts[cc] += 1
	}

	return len(counts)
}
