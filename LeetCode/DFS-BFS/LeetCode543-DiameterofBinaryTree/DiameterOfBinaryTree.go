package LeetCode543_DiameterofBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	if root == nil {
		return maxDia
	}

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		maxDia = max(maxDia, leftHeight+rightHeight)

		return 1 + max(leftHeight, rightHeight)
	}
	dfs(root)

	return maxDia
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
