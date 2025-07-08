package LeetCode150_EvaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		//取到的是数字就入栈
		value, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, value)
		} else {
			//不是数字的话，就取出栈中的倒数两个数，并且不需要入栈
			//根据符号，进行运算再将结果入栈
			nums1, nums2 := stack[len(stack)-2], stack[len(stack)-1]
			//缩减掉运算的数字
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, nums1+nums2)
			case "-":
				stack = append(stack, nums1-nums2)
			case "*":
				stack = append(stack, nums1*nums2)
			default:
				stack = append(stack, nums1/nums2)
			}
		}
	}
	return stack[0]
}
