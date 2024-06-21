package LeetCode367_ValidPerfectSquare

func isPerfectSquare(num int) bool {
	nums := 1
	for num > 0 {
		num -= nums
		nums += 2
	}
	return num == 0
}
