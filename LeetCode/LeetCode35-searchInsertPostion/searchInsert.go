package LeetCode35_searchInsertPostion


func searchInsert(nums []int, target int) int {
	right := len(nums)-1
	left := 0


	for left <= right{
		mid := (right-left)/2 + left
		if nums[mid] >target{
			right = mid -1
		}else if nums[mid] < target {
			left = mid +1
		}else {
			return mid
		}
	}
	return right+1
}