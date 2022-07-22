package LeetCode383_RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := [26]int{}

	for _, r := range magazine {
		cnt[r-'a']++
	}

	for _, v := range ransomNote {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
