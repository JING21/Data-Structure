package LeetCode76_MinimumWindowSubstring

import "math"

func minWindow(s string, t string) string {
	hs, ht := make(map[byte]int), make(map[byte]int)

	for i := range t {
		ht[t[i]]++
	}

	sLen := len(s)
	length := math.MaxInt32

	ansLeft, ansRight := -1, -1

	check := func() bool {
		for k, v := range ht {
			if hs[k] < v {
				return false
			}
		}
		return true
	}

	for left, right := 0, 0; right < sLen; right++ {
		if right < sLen && ht[s[right]] > 0 {
			hs[s[right]]++
		}
		for check() && left <= right {
			if right-left+1 < length {
				length = right - left + 1
				ansLeft, ansRight = left, left+length
			}

			if _, ok := ht[s[left]]; ok {
				hs[s[left]] -= 1
			}
			left++
		}
	}
	if ansLeft == -1 {
		return ""
	}
	return s[ansLeft:ansRight]
}
