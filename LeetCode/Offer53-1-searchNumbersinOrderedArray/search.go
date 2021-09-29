package Offer53_1_searchNumbersinOrderedArray

import "sort"

func search(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target + 1) - 1
	return rightmost - leftmost + 1
}



func search2(nums []int, target int) int {
	left := help(nums, 0, len(nums)-1, target)
	right := help(nums, left, len(nums)-1, target+1)
	return right-left
}

func help(nums []int, left, right, target int) int{
	for left <= right{
		mid := (left+right)/2
		if nums[mid] < target{
			left = mid+1
		}else {
			right = mid-1
		}
	}
	return left
}