package offer62_Josephus_problem


func lastRemaining(n int, m int) int {
	f := 0

	for i := 2; i <= n ; i++{
		f = (m + f) % i
	}

	return f
}