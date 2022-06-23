package offer42_MaxSubArray


func maxSubArray(nums []int) int {
	result := nums[0]
	for i:=1; i < len(nums); i++{
		nums[i] += max(nums[i-1],0)
		result = max(result, nums[i])
	}
	return result
}

func max(a,b int)int {
	if a>b{
		return a
	}
	return b
}


func min(a,b int)int {
	if a<b{
		return a
	}
	return b
}


