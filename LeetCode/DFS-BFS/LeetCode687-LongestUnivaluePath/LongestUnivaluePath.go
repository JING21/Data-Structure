package LeetCode687_LongestUnivaluePath

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		leftSame := 0
		rightSame := 0

		if node.Left != nil && node.Left.Val == node.Val {
			leftSame = +leftHeight + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			rightSame = +rightHeight + 1
		}
		result = max(result, leftSame+rightSame)
		return max(leftSame, rightSame)
	}

	dfs(root)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
