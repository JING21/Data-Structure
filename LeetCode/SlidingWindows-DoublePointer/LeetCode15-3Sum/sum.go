package LeetCode15_3Sum

import "sort"

// 1.先将nums排序,由小到大排序
// 2.将三个指针k,i,j分别设置,k为最左侧，i,j设置在数组(k,len(nums))两端
// 3.i,j交替向中间移动
// 4.nums[k] > 0时直接break,因为最小值都大于0，那三数之和不可能小于0
// 5.当k>0, 且nums[k] == nums[k-1],则跳过该元素，因为只计算一次
// 6.当i,j分别在数组索引(k, len(nums))两端，当i<j,计算三数之和 s=nums[k]+nums[i]+nums[j]
// s < 0时，i++,并且跳过所有重复的nums[i]
// s > 0时，j--，并且跳过所有重复的nums[j]
// s == 0时，i++,j--,并且跳过所有重复的nums[i], nums[j],并记录下k,i,j
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	//因为需要预留比k大的两个位置
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			//相等的情况
			if n1+n2+n3 == 0 {
				result = append(result, []int{n1, n2, n3})
				//跳过所有重复的nums[i]
				for left < right && nums[left] == n2 {
					left++
				}
				//跳过所有重复的nums[j]
				for left < right && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
