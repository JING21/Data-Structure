package LeetCode518_CoinChangeII

import (
	"encoding/json"
	"fmt"
	"os"
)

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))

	for i := range dp {
		dp[i] = make([]int, amount+1)
		//初始化左列，j=0时，一个硬币都不要，方案数量为1
		dp[i][0] = 1
	}

	//初始化最上行
	for j := coins[0]; j <= amount; j++ {
		dp[0][j] = dp[0][j-coins[0]]
	}

	//遍历物品
	for i := 1; i < len(coins); i++ {
		//遍历背包
		for j := 0; j <= amount; j++ {
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}
		}
	}
	return dp[len(coins)-1][amount]
}

func main() {
	// os.Args[0] 是程序名
	// os.Args[1:] 是实际参数
	if len(os.Args) < 4 { // 确保有至少三个参数
		fmt.Println("需要一个参数")
		fmt.Println("示例：", os.Args[0], os.Args[1], os.Args[2], "<参数1> <参数2> <参数3> ")
		return
	}
	arg0 := os.Args[1]
	arg1 := os.Args[2]
	fmt.Printf("参数1: %s\n参数2: %s\n", arg0, arg1)

	count, coin := dataProcess(arg0, arg1)

	maxValue := change(count, coin)
	fmt.Println(maxValue)
}

func dataProcess(arg0, arg1 string) (int, []int) {
	var count int
	err := json.Unmarshal([]byte(arg0), &count)
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}

	var coins []int
	err = json.Unmarshal([]byte(arg1), &coins)
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}

	return count, coins
}
