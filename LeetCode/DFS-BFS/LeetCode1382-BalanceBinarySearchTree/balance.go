package LeetCode1382_BalanceBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	result := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}

	var buildTree func(result []int, left, right int) *TreeNode
	buildTree = func(result []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := left + (right-left)>>1

		root := &TreeNode{result[mid], nil, nil}

		root.Left = buildTree(result, left, mid-1)
		root.Right = buildTree(result, mid+1, right)
		return root
	}
	dfs(root)

	return buildTree(result, 0, len(result)-1)
}
