package LeetCode309_BestTimetoBuyandSellStockWithCooldown

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	//分为三种情况：
	//1.dp[0][0] 表示持有一只股票，
	//2.dp[0][1] 表示不持有股票，但是处于冷冻期
	//3.dp[0][3] 表示不持有股票，但是不处于冷冻期

	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i < len(prices); i++ {
		//继续持有，保持前一天的状况，或者前一天没持有并且不在冷冻期，今天购入了股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		//在i天结束后不持有股票，并且处于冷冻期，说明了在i天进行了卖出交易,所以i-1天是持有股票的
		dp[i][1] = dp[i-1][0] + prices[i]
		//这个有两种情况1. i-1天也是这个不持有不处于冷冻期的情况，或者是i-1天的时候卖出了，但是处于冷冻期
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}
