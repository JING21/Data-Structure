package offer61_isStraight

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i:=0; i<4; i++{
		if nums[i] == 0{
			joker++
		}else if nums[i] == nums[i+1]{
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}