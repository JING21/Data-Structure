package offer21_adjustArray


func exchange(nums []int) []int {
	low, fast := 0, 0
	for fast < len(nums){
		if (nums[fast]&1 == 1){
			nums[fast], nums[low] = nums[low], nums[fast]
			low++
		}
		fast++
	}
	return nums
}


func exchange2(nums []int) []int{
	left, right := 0, len(nums)-1
	for left < right{
		if (nums[left]&1) == 0 && (nums[right]&1) == 1{
			nums[left], nums[right] = nums[right], nums[left]
		}
		for left < len(nums) && (nums[left]&1 != 0){
			left ++
		}
		for right >0 && (nums[right]&1 != 1){
			right --
		}
	}
	return nums
}


