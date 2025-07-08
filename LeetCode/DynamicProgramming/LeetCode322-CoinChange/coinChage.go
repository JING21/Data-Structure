package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

import "math"

func coinChange(coins []int, amount int) int {
	//dp表示从0-i取对应的硬币所组成的和为j的，硬币个数,而此时需要物品0，表示不取任何一个硬币的情况，所以需要len(coins)+1
	dp := make([][]int, len(coins)+1)

	for i := range dp {
		dp[i] = make([]int, amount+1)
		//初始化左列，当需要总数为0时，一个硬币也不需要
		dp[i][0] = 0
	}

	//初始化最上行为最大值，因为dp表示了需要的最少钱币数目，所以会取最小值，如果最后没有办法，则会保留最大值，return -1即可
	//所以将因为dp[0][x] = 最大值
	for j := 1; j <= amount; j++ {
		dp[0][j] = math.MaxInt32
	}

	//遍历物品
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			//表示不选择这个硬币
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				//如果选择这个硬币，就取选择这个硬币和不选择这个硬币的两者之间的最小值
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}
	if dp[len(coins)][amount] == math.MaxInt32 {
		return -1
	}
	return dp[len(coins)][amount]
}

func main() {
	// os.Args[0] 是程序名
	// os.Args[1:] 是实际参数
	if len(os.Args) < 2 { // 确保有至少一个参数
		fmt.Println("需要一个参数")
		fmt.Println("示例：", os.Args[2], "<参数1> <参数2>")
		return
	}

	arg2 := os.Args[2]

	fmt.Printf("参数1: %s\n参数2: %s\n", arg2)
	fmt.Println(parseWithStrings(arg2))
}

// 方法2：使用字符串操作（更高效）
func parseWithStrings(input string) []int {
	// 移除方括号和空格
	cleaned := strings.ReplaceAll(input, "[", "")
	cleaned = strings.ReplaceAll(cleaned, "]", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")
	fmt.Println(cleaned)
	// 分割字符串
	parts := strings.Split(cleaned, ",")
	fmt.Println(parts)
	fmt.Println(len(parts))
	nums := make([]int, len(parts))
	for i, part := range parts {
		// 移除单引号
		numStr := strings.Trim(part, "'")
		num, _ := strconv.Atoi(numStr)
		nums[i] = num
	}
	return nums
}
