package offer63_stockProfit

import "math"

func maxProfit2(prices []int) int {
	length:=len(prices)
	if length==0{return 0}
	dp:=make([][]int,length)
	for i:=0;i<length;i++{
		dp[i]=make([]int,2)
	}

	dp[0][0]=-prices[0]
	dp[0][1]=0
	for i:=1;i<length;i++{
		dp[i][0]=max(dp[i-1][0],-prices[i])
		dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max(a,b int)int {
	if a>b{
		return a
	}
	return b
}


func min(a,b int)int {
	if a<b{
		return a
	}
	return b
}


func maxProfit(prices []int) int{
	cost := math.MaxUint32
	profit := 0
	for _, v := range prices{
		cost = min(cost, v)
		profit = max(profit, v-cost)
	}
	return profit
}




