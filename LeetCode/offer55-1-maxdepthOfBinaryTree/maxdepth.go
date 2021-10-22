package offer55_1_maxdepthOfBinaryTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//DFS递归遍历
func maxDepth(root *TreeNode) int{
	if root == nil{
		return 0
	}

	return  max(maxDepth(root.Left), maxDepth(root.Right)) +1
}


func max(a,b int) int{
	if a > b {
		return a
	}
	return b
}