package LeetCode145_145_BinaryTreePostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if root == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)

	return result
}
