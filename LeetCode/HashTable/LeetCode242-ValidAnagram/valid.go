package LeetCode242_ValidAnagram

func isAnagram(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-rune('a')]++
	}

	for _, r := range t {
		record[r-rune('a')]--
	}

	return record == [26]int{}
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	exist := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if v, ok := exist[s[i]]; v >= 0 && ok {
			exist[s[i]] = v + 1
		} else {
			exist[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if v, ok := exist[t[i]]; v >= 1 && ok {
			exist[t[i]] = v - 1
		} else {
			return false
		}
	}

	return true
}
