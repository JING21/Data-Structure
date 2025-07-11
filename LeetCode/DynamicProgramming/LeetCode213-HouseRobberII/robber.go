package LeetCode213_HouseRobberII

func rob(nums []int) int {
	//因为有环，所以包含三个情况
	//1.首尾都不抢
	//2.抢首不抢尾
	//3.抢尾不抢首

	//只有一间的情况
	if len(nums) <= 1 {
		return robWithoutCircle(nums)
	}

	result1 := robWithoutCircle(nums[:len(nums)-1])
	result2 := robWithoutCircle(nums[1:])

	return max(result1, result2)

}

func robWithoutCircle(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	//dp表示偷到nums[i]家能够得到的最大金额,0表示第一家
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		//dp[i]应该取偷到前一家的金额和前前一家金额+当前这一家金额之和,两者的最大值
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}
