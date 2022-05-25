package LeetCode107_BinaryTreeLevelOrderTraversalII_bottomToTop

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	for k := 0; k < len(result)/2; k++ {
		result[k], result[len(result)-k-1] = result[len(result)-k-1], result[k]
	}

	return result
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
