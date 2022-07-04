package LeetCode18_4Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]

		if i > 0 && n1 == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1

			for left < right {
				n3 := nums[left]
				n4 := nums[right]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{n1, n2, n3, n4})
					for left < right && n3 == nums[left+1] {
						left++
					}
					for left < right && n4 == nums[right-1] {
						right--
					}
					right--
					left++
				}
			}
		}
	}
	return result
}
