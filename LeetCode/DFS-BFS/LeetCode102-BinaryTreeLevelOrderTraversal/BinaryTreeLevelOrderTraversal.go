package LeetCode102_BinaryTreeLeverlOrderTraversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func levelOrder2(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	currentLevel := []*TreeNode{root}

	for i := 0; len(currentLevel) > 0; i++ {
		result = append(result, []int{})
		nextLevel := []*TreeNode{}
		for j := 0; j < len(currentLevel); j++ {
			node := currentLevel[j]
			result[i] = append(result[i], node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
	}
	return result
}
