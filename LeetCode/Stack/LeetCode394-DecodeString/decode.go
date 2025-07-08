package LeetCode394_DecodeString

func decodeString(s string) string {
	var stack []byte

	//遍历入栈
	for i := 0; i < len(s); i++ {
		//不遇到']'之前，都入栈
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			//辅助栈
			tmp := []byte{}
			//遇到之后取出前序字符，直到'['为止
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			//弹出了左括号,'['
			stack = stack[:len(stack)-1]

			//取出括号前数字,因为数字不一定是个位数所以需要累加计算
			//需要乘上对应的个位，十位，百位，用base表示
			num := 0
			base := 1

			//判断是否是数字
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				num += int(stack[len(stack)-1]-'0') * base
				//计算下一位
				base *= 10
				//吐出数字
				stack = stack[:len(stack)-1]
			}

			//反转tmp，因为是后面的字符先传出来
			for j := 0; j < len(tmp)/2; j++ {
				tmp[j], tmp[len(tmp)-1-j] = tmp[len(tmp)-1-j], tmp[j]
			}
			//将正确打印的压入栈中
			for j := 0; j < num; j++ {
				stack = append(stack, tmp...)
			}
		}
	}
	return string(stack)
}
