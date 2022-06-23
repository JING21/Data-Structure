package offer10_2_frogJumping


func numWays(n int) int {
	const mod int = 1e9+7
	if n <= 1{
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b, sum := 0, 1, 2
	for i := 3; i <=n; i++{
		a = b
		b = sum
		sum = (a+b) % mod
	}
	return sum
}