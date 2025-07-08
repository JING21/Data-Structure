package LeetCode300_LongestIncreasingSubsequence

func lengthOfLIS(nums []int) int {
	//dp[i]表示以i结尾的取nums[i]这个元素时，子序列的长度
	dp := make([]int, len(nums))

	//初始化dp
	for i := range dp {
		dp[i] = 1
	}
	//初始化结果
	ans := dp[0]

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			//如果后一位大于前一位的话，则是递增的子序列
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
