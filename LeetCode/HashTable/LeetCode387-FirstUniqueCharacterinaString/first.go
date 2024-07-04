package LeetCode387_FirstUniqueCharacterinaString

func firstUniqChar(s string) int {
	set := make(map[byte]int)

	for i := range s {
		set[s[i]]++
	}

	for i := range s {
		if set[s[i]] == 1 {
			return i
		}
	}

	return -1
}
