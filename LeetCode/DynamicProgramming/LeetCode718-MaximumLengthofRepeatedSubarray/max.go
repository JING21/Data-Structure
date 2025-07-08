package LeetCode718_MaximumLengthofRepeatedSubarray

func findLength(nums1 []int, nums2 []int) int {
	//dp表示以i-1，j-1结尾的nums1,,nums2的最长重复子数组的值
	//为什么是i-1和j-1呢，因为dp[i][j]只能由dp[i-1][j-1]推导出来
	//当头两个字母相同时，dp[0][0]=1
	//并且第一行或者第一列需要单独初始化，所以填入一个空值，可以是第一行第一列都是0，不需要额外初始化

	m, n := len(nums1), len(nums2)
	result := 0

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result
}
