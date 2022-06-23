package offer53_2_searchMissingNumber

func missingNumber(nums []int) int {
	n:= len(nums)
	left, right := 0, n-1
	for left <= right{
		mid := (left+right)/2
		if nums[mid] == mid{
			left = mid+1
		}else {
			right = mid-1
		}
	}
	return left
}
