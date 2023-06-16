package main

import "sort"

func unique_chars(s string) []string {
	counts := map[string]int{}
	for _, c := range s {
		cc := string(c)
		counts[cc] += 1
	}

	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
