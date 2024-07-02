package LeetCode714_BestTimetoBuyandSellStockwithTransactionFee

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	//包含两种情况
	//dp[i][0]表示没持有股票
	//dp[i][1]表示持有股票
	dp := make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		//不持有包含两种情况
		//1.继续前一天没持有的状态
		//2.前一天持有，今天卖出了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		//持有也包含两种情况
		//1.前一天持有，今天继续持有
		//2.前一天没持有，今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}
