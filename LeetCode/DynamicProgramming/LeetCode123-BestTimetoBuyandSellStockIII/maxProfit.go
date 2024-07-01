package LeetCode123_BestTimetoBuyandSellStockIII

func maxProfit(prices []int) int {
	n := len(prices)

	if n == 0 {
		return 0
	}

	dp := make([][5]int, n)
	//分为五个状态，0表示没有操作，1表示第一次持有股票，2表示第一次不持有股票，3表示第二次持有股票
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < n; i++ {
		//未操作
		dp[i][0] = dp[i-1][0]
		//两种情况，1.未操作延续i-1天，未买入股票，但是持有 2.当天买入股票
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		//两种情况，1.未操作延续i-1天，未持有股票，但是持有 2.当天卖出股票股票
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])
		//同理如dp[i][1]
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		//同理如dp[i][2]
		dp[i][4] = max(dp[i-1][3]+prices[i], dp[i-1][4])
	}

	return dp[n-1][4]

}
