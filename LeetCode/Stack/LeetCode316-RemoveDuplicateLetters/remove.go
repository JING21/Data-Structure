package LeetCode316_RemoveDuplicateLetters

func removeDuplicateLetters(s string) string {
	//每个字母出现次数
	count := [26]int{}
	//是否已在栈中
	exist := [26]bool{}
	//单调栈
	stack := []rune{}

	//统计每个字母出现的次数
	for _, c := range s {
		count[c-'a']++
	}

	//遍历字符串
	for _, c := range s {
		//栈中已存在，跳过这个字母
		//并且减少计数
		if exist[c-'a'] {
			count[c-'a']--
			continue
		}

		//出栈的条件
		//1.栈不为空
		//2.栈里所有元素都大于当前元素C
		//3.栈顶的元素后续还有，当前的即可舍弃
		for len(stack) > 0 && stack[len(stack)-1] > c && count[stack[len(stack)-1]-'a'] > 0 {
			//将是否存在栈内设置为false
			exist[stack[len(stack)-1]-'a'] = false
			//删除栈顶元素
			stack = stack[:len(stack)-1]
		}

		//入栈操作
		stack = append(stack, c)
		//减少次数
		count[c-'a']--
		//设置栈内有
		exist[c-'a'] = true
	}
	return string(stack)
}
