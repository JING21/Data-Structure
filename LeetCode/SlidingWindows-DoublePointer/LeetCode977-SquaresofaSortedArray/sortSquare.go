package LeetCode977_SquaresofaSortedArray

func sortedSquares(nums []int) []int {
	len := len(nums)

	if len == 0 {
		return nil
	}

	slow, fast := 0, len-1
	result := make([]int, len)

	for pos := len - 1; pos >= 0; pos-- {
		if v, w := nums[slow]*nums[slow], nums[fast]*nums[fast]; v > w {
			result[pos] = v
			slow++
		} else {
			result[pos] = w
			fast--
		}

	}
	return result
}
