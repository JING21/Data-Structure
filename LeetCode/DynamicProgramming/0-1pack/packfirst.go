package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pack(kind, bagweight int, weight, value []int) int {
	dp := make([][]int, kind)

	//初始化每个物品对应的每个容量的背包的价值数组
	for i := range dp {
		dp[i] = make([]int, bagweight+1)
	}

	//当背包容量 < 0号物品的重量时，表示背包连0号都装不下,那么dp对应的value应该是0，dp[0][j] = 0
	//当背包容量 > 0号物品的重量时，表示背包可以装下0号物品，那么dp对应的value应该是0号物品的价值，dp[0][j] = value[0]
	for j := weight[0]; j <= bagweight; j++ {
		dp[0][j] = value[0]
	}

	for i := 1; i < kind; i++ { //遍历物品种类
		for j := 0; j <= bagweight; j++ { //遍历背包容量
			if j < weight[i] { //如果装不下这个物品，那么只能不装这个物品，使用上个物品的值
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[kind-1][bagweight]
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
	arg2 := os.Args[3]
	fmt.Printf("参数1: %s\n参数2: %s\n参数3: %s\n", arg0, arg1, arg2)

	kind, bagWeight, weight, value := dataProcess(arg0, arg1, arg2)

	maxValue := pack(kind, bagWeight, weight, value)
	fmt.Println(maxValue)
}

func dataProcess(arg1, arg2, arg3 string) (kind, bagWeight int, weight, value []int) {
	data1 := removeBrackets(arg1)
	data2 := removeBrackets(arg2)
	data3 := removeBrackets(arg3)

	kind, _ = strconv.Atoi(data1[0])
	bagWeight, _ = strconv.Atoi(data1[1])

	for _, v := range data2 {
		i, _ := strconv.Atoi(v)
		weight = append(weight, i)
	}

	for _, v := range data3 {
		i, _ := strconv.Atoi(v)
		value = append(value, i)
	}
	fmt.Println(kind)
	fmt.Println(bagWeight)
	fmt.Println(weight)
	fmt.Println(value)
	return kind, bagWeight, weight, value
}

func removeBrackets(arg string) []string {
	arg = arg[1 : len(arg)-1]
	sd := strings.Split(arg, ",")
	return sd
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
