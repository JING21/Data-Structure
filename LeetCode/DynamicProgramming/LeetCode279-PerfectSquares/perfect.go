package LeetCode279_PerfectSquares

import "math"

func numSquares(n int) int {
	//dp表示满足和为n的，完全平方数的最小个数
	//其中i应该取值为各个完全平方数，0,1,4,9,16, i*i < n才满足条件
	//由于可以反复取完全平方数，所以是完全背包问题
	x := integerSquareRoot(n)
	dp := make([][]int, x)

	for i := range dp {
		dp[i] = make([]int, n+1)
		//因为要取最小值，所以要初始化为最大值，避免初始化为空，干扰
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	//初始化首位,因为n>1
	dp[0][0] = 0
	//初始化最上行,只有1的情况
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	//初始化第一列，当n=0的情况
	for i := 0; i < x; i++ {
		dp[i][0] = 0
	}

	//dp过程，从1开始处理，因为1*1可以组成任意正数和
	for i := 1; i < x; i++ {
		weight := (i + 1) * (i + 1)
		for j := 0; j <= n; j++ {
			//当前数+了之后，大于J了，或者前一个也是无值可取的情况，那当前还是继续继承
			if j < weight || dp[i][j-weight] == math.MaxInt32 {
				dp[i][j] = dp[i-1][j]
			} else {
				//完全背包所以是dp[i][]不是dp[i-1]
				dp[i][j] = min(dp[i-1][j], dp[i][j-weight]+1)
			}
		}
	}
	return dp[x-1][n]
}

func integerSquareRoot(n int) int {
	if n < 0 {
		return 0 // 负数返回0，或根据需求返回错误
	}
	if n == 0 {
		return 0
	}
	k := int(math.Sqrt(float64(n)))
	//判断是否为完全平方数
	if (k + 1) <= n/(k+1) {
		return k + 1
	}
	return k
}
