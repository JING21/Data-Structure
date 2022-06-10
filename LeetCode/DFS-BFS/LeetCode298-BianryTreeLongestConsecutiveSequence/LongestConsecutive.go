package LeetCode298_BianryTreeLongestConsecutiveSequence

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxLength int

func longestConsecutive(root *TreeNode) int {
	maxLength := 0
	var dfs func(node, parent *TreeNode, length int)
	dfs = func(node, parent *TreeNode, length int) {
		if node == nil {
			return
		}
		if parent != nil && node.Val == parent.Val+1 {
			length = length + 1
		} else {
			length = 1
		}
		maxLength = max(maxLength, length)
		dfs(node.Left, node, length)
		dfs(node.Right, node, length)
	}
	dfs(root, nil, 0)

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
