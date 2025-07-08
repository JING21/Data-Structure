package LeetCode739_DailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	days := len(temperatures)
	ans := make([]int, days)
	//栈入队的是温度的index序号日期而不是对应的温度
	stack := []int{}

	for i := 0; i < days; i++ {
		temperature := temperatures[i]
		//判断条件是，栈不为空，并且当且温度大于栈最上方的温度
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		//入栈
		stack = append(stack, i)
	}

	return ans
}
