package LeetCode152_MaximumProductSubarray

func maxProduct(nums []int) int {
	//需要维护两个数组
	//1.是以i结尾的0-i的最大值(正数)
	//2.是以i结尾的0-i的最小值（负数）
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	maxResult := nums[0]

	for i := 1; i < len(nums); i++ {
		//以i结尾的nums[i]的最大的子数组乘积包含三种情况，
		//当nums[i]为负数时，希望前面连续子集乘积为尽可能小的负数
		//当nums[i]为正数时，希望前面连续子集乘积为尽可能大的正数
		//或者是nums[i]为正数时，dp[i-1]为负数时，就是nums[i]最大
		dpMax[i] = max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
		dpMin[i] = min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])

		maxResult = max(maxResult, dpMax[i])
	}

	return maxResult
}
