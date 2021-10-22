package offer55_2_balanceTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//DFS后序遍历
func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return height(root) > 0
}

func height(root *TreeNode) int{
	if root == nil{
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1{
		return -1
	}

	return max(leftHeight, rightHeight)+1
}



func max(a,b int) int{
	if a > b {
		return a
	}
	return b
}


func abs(a int) int{
	if a < 0 {
		return -1 * a
	}
	return a
}