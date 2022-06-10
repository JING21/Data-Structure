package LeetCode98_ValidateBinarySearchTree_

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	result := []*TreeNode{}

	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node)
		dfs(node.Right)
	}
	dfs(root)

	for i := 1; i < len(result); i++ {
		if result[i].Val <= result[i-1].Val {
			return false
		}
	}

	return true
}

func isValidBST2(root *TreeNode) bool {
	var helper func(node *TreeNode, lower, upper int) bool
	helper = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}

		if node.Val <= lower || node.Val >= upper {
			return false
		}

		return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, upper)
	}

	return helper(root, math.MinInt64, math.MaxInt64)
}

func isValidBST3(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
