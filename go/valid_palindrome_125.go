package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	var spaceRegex = regexp.MustCompile(` `)

	s_normalized := nonAlphanumericRegex.ReplaceAllString(strings.ToLower(s), "")
	s_normalized = spaceRegex.ReplaceAllString(s_normalized, "")

	length := len(s_normalized)
	is_palindrome := true
	for i := range s_normalized {
		left := string(s_normalized[i])
		right := string(s_normalized[(length-1)-i])

		if i > (length-1)-i {
			break
		}

		if left != right {
			is_palindrome = false
			break
		}
	}

	return is_palindrome
}
