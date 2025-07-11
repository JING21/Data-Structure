package LeetCode198_HouseRobber

func rob(nums []int) int {
	//dp表示偷到nums[i]家能够得到的最大金额
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	dp[0] = 0

	for i := 2; i <= len(nums); i++ {
		//dp[i]应该取偷到前一家的金额和前前一家金额+当前这一家金额之和,两者的最大值
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}
