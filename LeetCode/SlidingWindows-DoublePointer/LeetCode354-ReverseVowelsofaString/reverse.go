package LeetCode354_ReverseVowelsofaString

import "strings"

func reverseVowels(s string) string {
	t := []byte(s)
	n := len(s)
	l, r := 0, n-1

	for l < r {
		for l < n && !strings.Contains("aeiouAEIOU", string(t[l])) {
			l++
		}

		for r > 0 && !strings.Contains("aeiouAEIOU", string(t[r])) {
			r--
		}
		if l < r {
			t[l], t[r] = t[r], t[l]
			l++
			r--
		}
	}
	return string(t)
}
