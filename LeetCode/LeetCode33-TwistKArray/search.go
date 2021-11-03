package LeetCode33_TwistKArray


func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right{
		mid := (right-left)/2 + left

		if nums[mid] == target{
			return mid
		}

		if nums[left] <= nums[mid]{
			if nums[left] <= target && nums[mid] > target{
				right = mid - 1
			}else {
				left = mid + 1
			}
		}else {
			if nums[mid] < target && nums[right] >= target{
				left = mid + 1
			}else {
				right = mid - 1
			}

		}
	}
	return -1
}