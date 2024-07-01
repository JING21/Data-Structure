package LeetCode122_BestTimetoBuyandSellStockII

func maxProfit(prices []int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}

	//j=0:持有现金
	//j=1:持有股票
	//状态转移是0-1-0-1-0-1-0
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		//分为两种情况前一天不持有股票，今天仍继续不持有，或者前一天持有股票，今天不持有了卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		//前一天持有股票，今天仍继续持有，或者前一天不持有股票，今天买入了
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}
