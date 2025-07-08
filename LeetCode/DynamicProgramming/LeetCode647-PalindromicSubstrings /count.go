package LeetCode647_PalindromicSubstrings_

func countSubstrings(s string) int {
	result := 0
	//回文串的情况分为3种
	//1.只有一个字母，自身就是回文串
	//2.如果是两个字母，则首尾字母相等
	//3.如果大于2个字母，那么去掉首尾两个字母后，应该仍然是回文字母
	//dp[i][j]表示从i开始到j的字符串是否为回文串

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	//当dp[i][j]中，s[i] = s[j]时，只要判断dp[i+1][j-1]是否回文即可
	//所以不能按正常顺序遍历，要按照从下到上，从左到右进行遍历
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				//情况1和情况2
				if j-i <= 1 {
					result++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					result++
					dp[i][j] = true
				}
			}
		}
	}
	return result
}
