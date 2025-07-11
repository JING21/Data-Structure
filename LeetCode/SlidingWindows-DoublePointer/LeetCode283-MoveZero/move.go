package LeetCode283_MoveZero

func moveZeroes(nums []int) {
	slow, fast, len := 0, 0, len(nums)

	//快慢指针，参考快速排序，把不等于0的都一直放在右边即可
	for fast < len {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
