package offer03_repeatedNumbers

func findRepeatNumber(nums []int) int {
	hash := make(map[int]bool, len(nums))
	var b int
	for i := 0; i<len(nums); i++{
		if hash[nums[i]]{
			b = nums[i]
			break
		}
		hash[nums[i]] = true
	}
	return b
}

func findRepeatNumber2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			swap(nums[i], nums[nums[i]])
		}
	}
	return -1
}


func swap(a,b int) {
	var temp int
	temp = a
	a = b
	b = temp
}