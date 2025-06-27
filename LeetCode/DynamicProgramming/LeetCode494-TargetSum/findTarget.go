package LeetCode494_TargetSum

import "math"

func findTargetSumWays(nums []int, target int) int {
	//数组和
	sum := 0
	for _, v := range nums {
		sum += v
	}

	//复数和假设为neg,则正数和应该是sum-neg,然后(sum-neg)-neg=target(目标和)，所以neg=(sum-target)/2,positive=(sum+target)/2

	bagSize := (sum + target) / 2

	//如果target大于整个数组和，那么全部+或者全部-也不能满足
	//并且neg必须为非负偶数才能满足，所以(sum-target)/2需要是非负偶数，所以其为奇数时，直接不满足
	if math.Abs(float64(target)) > math.Abs(float64(sum)) {
		return 0
	}
	if (sum-target)%2 == 1 {
		return 0
	}

	//dp表示对应的方案数目，表示数组中选取前i个数字，是元素之和为j的方案数目
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, bagSize+1)
	}

	//初始化第上行，dp[0][第一个数]对应的j，只有唯一一种方法，就是放入nums[0]实现j=nums[0],其他都无法实现，因为只使用nums[0]这一个数
	if nums[0] <= bagSize {
		dp[0][nums[0]] = 1
	}

	//初始化第左列
	//当nums数组为0时，元素和只能是0，只有一种方案，不放，没有数
	dp[0][0] = 1

	//特殊情况，当数组中包含0的数字时，满足j=0的情况就变多了，比如说有物品0是0，那装满0的情况就是，1.不放 2放入0两种，
	//当物品0，1都为0时，那满足的情况的方法就有，1.不放 2.放物品0 3.放物品1 4.放物品0，1
	//当物品0，1，2都为0时，那满足的情况就是1.不放 2.放物品0 3.放物品1 4.放物品2 5.放物品0,1 6.放物品0,2 7，放物品1，2 8.放物品0，1，2
	var numZero float64
	for i, v := range nums {
		if v == 0 {
			numZero++
		}
		dp[i][0] = int(math.Pow(2, numZero))
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= bagSize; j++ {
			if j < nums[i] { //当nums[i] > j时，这时候nums[i]一定不能取，所以是dp[i - 1][j]种方案数
				dp[i][j] = dp[i-1][j]
			} else { //nums[i] <= j时，num[i]可取可不取，因此方案数是dp[i - 1][j] + dp[i - 1][j - nums[i]]
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][bagSize]
}
