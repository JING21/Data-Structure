package LeetCode58_LastLengthOfWord

func lengthOfLastWord(s string) int {
	length := len(s)
	index := length-1
	ans := 0
	for s[index] == ' '{
		index--
	}

	for index >= 0 && s[index] != ' '{
		index --
		ans++
	}

	return ans
}
