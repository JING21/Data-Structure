package LeetCode26_RemoveDuplicatesfromSortedArray

func removeDuplicates(nums []int) int {
	len := len(nums)

	if len == 0 {
		return 0
	}
	slow := 1

	for fast := 1; fast < len; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
