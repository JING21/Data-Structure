package offer28_isSymmetric

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isPair(root.Left, root.Right)
}

func isPair(Left, right *TreeNode) bool {
	if Left == nil && right == nil{
		return true
	}
	if Left == nil || right == nil{
		return false
	}
	if Left.Val != right.Val{
		return false
	}
	return isPair(Left.Left, right.Right) && isPair(Left.Right, right.Left)
}