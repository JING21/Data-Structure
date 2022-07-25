package LeetCode202_HappyNumber

func isHappy(n int) bool {
	m := make(map[int]bool)

	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}

	return n == 1
}

func isHappy2(n int) bool {
	slow, fast := n, getSum(n)

	for slow != fast && fast != 1 {
		slow = getSum(slow)
		fast = getSum(fast)
		fast = getSum(fast)
	}

	return fast == 1
}

func getSum(n int) int {
	sum := 0

	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}

	return sum
}
