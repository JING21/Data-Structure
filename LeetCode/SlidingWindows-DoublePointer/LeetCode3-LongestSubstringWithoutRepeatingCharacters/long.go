package LeetCode3_LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	count := map[byte]int{}
	left := 0
	result := 0

	for right := range s {
		count[s[right]]++
		//如果窗口内已经包含了c，那么需要移除窗口内的再加入
		for count[s[right]] > 1 {
			count[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
