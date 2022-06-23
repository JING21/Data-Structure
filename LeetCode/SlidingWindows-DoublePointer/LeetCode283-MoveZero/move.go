package LeetCode283_MoveZero

func moveZeroes(nums []int) {
	slow, fast, len := 0, 0, len(nums)

	for fast < len {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
