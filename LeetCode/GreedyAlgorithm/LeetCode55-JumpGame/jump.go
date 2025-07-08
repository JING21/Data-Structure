package LeetCode55_JumpGame

func canJump(nums []int) bool {
	//转换为覆盖问题
	cover := 0
	n := len(nums) - 1

	for i := 0; i <= cover; i++ {
		cover = max(nums[i]+i, cover)
		if cover >= n {
			return true
		}
	}
	return false
}
