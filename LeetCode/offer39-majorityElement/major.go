package offer39_majorityElement

//摩尔投票算法
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, v := range nums{
		if count == 0{
			candidate = v
		}
		if v == candidate{
			count ++
		}else {
			count--
		}
	}
	return candidate
}

