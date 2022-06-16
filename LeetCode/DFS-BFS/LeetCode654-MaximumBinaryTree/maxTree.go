package LeetCode654_MaximumBinaryTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var dfs func(nums []int) *TreeNode
	dfs = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}

		maxIndex := 0
		maxNum := math.MinInt64
		root := &TreeNode{}
		for k, v := range nums {
			if v > maxNum {
				maxNum = v
				maxIndex = k
			}
		}

		root.Val = maxNum
		root.Left = dfs(nums[:maxIndex])
		root.Right = dfs(nums[maxIndex+1:])

		return root
	}
	return dfs(nums)
}
