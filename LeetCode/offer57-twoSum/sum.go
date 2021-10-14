package offer57_twoSum

func twoSum(nums []int, target int) []int {
	 left, right := 0, len(nums)-1
	 for left < right{
	 	s := nums[left] + nums[right]
	 	if s < target{
	 		left ++
		} else if s > target {
			right --
		} else {
			a := []int{ nums[left], nums[right]}
			return a
		}
	 }
	 return []int{}
}