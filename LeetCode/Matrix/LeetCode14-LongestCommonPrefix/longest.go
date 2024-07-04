package LeetCode14_LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	//
	s0 := strs[0]
	for j := range s0 {
		for _, v := range strs {
			if j == len(v) || v[j] != s0[j] {
				return s0[:j]
			}
		}
	}

	return s0
}
