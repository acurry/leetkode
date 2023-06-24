package main

func isSubsequence(s string, t string) bool {
	s_index := 0
	if len(s) == 0 {
		return true
	}

	for _, c := range t {
		if s_index == len(s) {
			return true
		}
		if string(s[s_index]) == string(c) {
			s_index += 1
		}
	}

	return len(s[s_index:]) == 0
}
