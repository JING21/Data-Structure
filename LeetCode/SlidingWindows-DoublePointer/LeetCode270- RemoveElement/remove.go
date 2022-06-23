package LeetCode270__RemoveElement

func removeElement(nums []int, val int) int {
	len := len(nums)
	res := 0

	if len == 0 {
		return 0
	}

	for i := 0; i < len; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
