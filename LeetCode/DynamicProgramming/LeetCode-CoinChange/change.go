package LeetCode_CoinChange

import "math"

func coinChange(coins []int, amount int) int {
	//dp表示从0-i取对应的硬币所组成的和为j的，硬币个数，此时硬币序列应该为硬币数目+1，因为需要考虑不拿硬币的情况
	dp := make([][]int, len(coins)+1)

	for i := range dp {
		dp[i] = make([]int, amount+1)
		//初始化左列，当需要总数为0时，一个硬币也不需要
		dp[i][0] = 0
	}

	//初始化最上行为最大值，因为dp表示了需要的最少钱币数目，所以会取最小值，如果最后没有办法，则会保留最大值，return -1即可
	//所以将因为dp[0][x] = 最大值
	for j := 1; j <= amount; j++ {
		dp[0][j] = math.MaxInt32
	}

	//遍历物品
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			//表示不选择这个硬币
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				//如果选择这个硬币，就取选择这个硬币和不选择这个硬币的两者之间的最小值
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}
	if dp[len(coins)][amount] == math.MaxInt32 {
		return -1
	}
	return dp[len(coins)][amount]
}
