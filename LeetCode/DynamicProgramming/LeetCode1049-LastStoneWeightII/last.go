package LeetCode1049_LastStoneWeightII

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2

	dp := make([][]int, len(stones))

	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for j := stones[0]; j <= target; j++ {
		dp[0][j] = stones[0]
	}

	for i := 1; i < len(stones); i++ {
		for j := 0; j <= target; j++ {
			if j < stones[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			}
		}
	}

	return sum - 2*dp[len(stones)-1][target]
}

