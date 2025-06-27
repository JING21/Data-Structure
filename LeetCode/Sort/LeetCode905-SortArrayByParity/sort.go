package LeetCode905_SortArrayByParity

func sortArrayByParity(nums []int) []int {
	for l, r := 0, len(nums)-1; l < r; {
		for l < r && nums[l]%2 == 0 {
			l++
		}
		for r > l && nums[r]%2 == 1 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
