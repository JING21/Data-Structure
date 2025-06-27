package LeetCode416_PartitionEqualSubsetSum

import (
	"encoding/json"
	"fmt"
	"os"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([][]int, len(nums))

	//初始化dp[i]
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	//初始化dp[0][j]
	for j := nums[0]; j <= target; j++ {
		dp[0][j] = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
		}
	}
	return dp[len(nums)-1][target] == target
}

func main() {
	// os.Args[0] 是程序名
	// os.Args[1:] 是实际参数
	if len(os.Args) < 2 { // 确保有至少一个参数
		fmt.Println("需要一个参数")
		fmt.Println("示例：", os.Args[1], "<参数1> <参数2>")
		return
	}
	arg0 := os.Args[0]
	arg1 := os.Args[1]

	fmt.Printf("参数1: %s\n参数2: %s\n", arg0, arg1)
	fmt.Println(parseWithStrings(arg1))
}

func parseWithStrings(arg string) []int {
	var num []int

	err := json.Unmarshal([]byte(arg), &num)
	if err != nil {
		fmt.Println(err)
	}

	return num
}

func max(x, y int) int {

	if x > y {
		return x
	}

	return y
}