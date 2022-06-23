package offer32_BFS_PrintBinaryTree

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
 type TreeNode struct {
    Val 	int
    Left 	*TreeNode
    Right 	*TreeNode
 }

func levelOrder(root *TreeNode) []int {
	if root == nil{
		return nil
	}

	result := make([]int, 0)

	queue := make([]*TreeNode, 0)

	queue = append(queue, root)

	for len(queue) > 0{
		tmp := make([]*TreeNode, 0)
		for _, node := range queue {
			result = append(result, node.Val)
			if node.Left != nil{
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil{
				tmp = append(tmp, node.Right)
			}
			queue = tmp
		}
	}
	return result
}