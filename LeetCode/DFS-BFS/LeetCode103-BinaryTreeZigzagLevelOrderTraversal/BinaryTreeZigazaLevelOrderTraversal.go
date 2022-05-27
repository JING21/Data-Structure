package LeetCode103_BinaryTreeZigzagLevelOrderTraversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 1 {
			for i, j := 0, len(vals); i < j/2; i++ {
				vals[i], vals[j-1-i] = vals[j-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return
}
