package LeetCode168_ExcelSheetColumnTitle

func convertToTitle(columnNumber int) string {
	s := ""
	for columnNumber > 0 {
		columnNumber--
		rest := columnNumber % 26
		s = string('A'+rest) + s
		columnNumber /= 26
	}

	return s
}
