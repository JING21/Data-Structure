package LeetCode1768_MergeStringsAlternately

func mergeAlternately(word1 string, word2 string) string {
	n, m := len(word1), len(word2)

	result := make([]byte, n+m)

	for i := 0; i < n || i < m; i++ {
		if i < n {
			result = append(result, word1[i])
		}
		if i < m {
			result = append(result, word2[i])
		}
	}
	return string(result)
}
