package LeetCode268_MissingNumber

func missingNumber(nums []int) int {
	n := len(nums)

	total := (1 + n) * n / 2

	for _, v := range nums {
		total -= v
	}

	return total
}
