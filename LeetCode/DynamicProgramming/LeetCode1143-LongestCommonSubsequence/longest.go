package LeetCode1143_LongestCommonSubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[i][j]表示从text1[0,i-1]与text2[0,j-1]中的最长序列是多少
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
