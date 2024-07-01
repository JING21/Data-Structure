package LeetCode121_BestTimetoBuyandSellStock

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	minPrice := prices[0]
	//表示第i天的最大收益
	dp := make([]int, n)

	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		//表示不操作买卖或者买卖两种情况
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}

	return dp[n-1]
}
