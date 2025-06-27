package LeetCode53_MaximumSubarray

func maxSubArray(nums []int) int {
	//i表示nums索引序号，dp[i]表示以nums[i]结尾的子序列和的最大值，
	dp := make([]int, len(nums))
	//最开始的最大值是nums[0]
	mx := nums[0]
	//初始化dp，以第一个数结尾的和的最大子序列应该就是nums[0]自身
	dp[0] = nums[0]

	//遍历数组
	for i := 1; i < len(nums); i++ {
		//会有两种情况，1.加上前面的和更大 2.不加前面的和
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		mx = max(dp[i], mx)
	}

	return mx
}
