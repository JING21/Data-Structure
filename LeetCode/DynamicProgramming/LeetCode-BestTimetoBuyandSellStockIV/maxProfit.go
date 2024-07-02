package LeetCode_BestTimetoBuyandSellStockIV

func maxProfit(k int, prices []int) int {
	n := len(prices)

	if n == 0 {
		return 0
	}

	k = min(k, n/2)
	//buy[i][j]表示持有一只股票，并且恰好进行j笔交易
	buy := make([][]int, n)
	//sell[i][j]表示不持有股票，并且恰好进行j笔交易
	sell := make([][]int, n)

	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}

	buy[0][0] = -prices[0]
	sell[0][0] = 0

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j < k; j++ {
			//比较如果是第i天买入的，那么i-1天不持有股票即sell[i-1][j]再减去买入的prices[i]
			//如果不是i天买入的,那么就是i-1天就持有股票的情况，即buy[i-1][j]
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			//比较如果是第i天卖出的，那么i-1天持有股票即buy[i-1][j-1]再加上卖出的prices[i]
			//如果不是i天买入的,那么就是i-1天就持有股票的情况，即buy[i-1][j]
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}

	return max(sell[n-1]...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func maxProfit2(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	status := make([]int, (2*k+1)*len(prices))
	for i := range dp {
		dp[i] = status[:2*k+1]
		status = status[2*k+1:]
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			//j为奇数是是买，为偶数时为卖
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}
