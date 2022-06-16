package LeetCode111_MininumDepthofBinaryTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	mindepth := math.MaxInt64
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil {
		mindepth = min(minDepth(root.Left), mindepth)
	}

	if root.Right != nil {
		mindepth = min(minDepth(root.Right), mindepth)
	}

	return mindepth + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
