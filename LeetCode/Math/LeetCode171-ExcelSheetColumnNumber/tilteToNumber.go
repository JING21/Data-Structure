package LeetCode171_ExcelSheetColumnNumber

func titleToNumber(columnTitle string) int {
	ans := 0
	for _, v := range columnTitle {
		num := v - 'A' + 1
		ans = ans*26 + int(num)
	}

	return ans
}

func titleToNumber2(columnTitle string) int {
	num := 0
	pow := 1

	for i := len(columnTitle) - 1; i >= 0; i-- {
		num += pow * (int(columnTitle[i]) - int('A') + 1)
		pow *= 26
	}
	return num
}
