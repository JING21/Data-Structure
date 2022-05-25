package LeetCode429_N_aryTreeLevelOrderTraversal

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	currentLevel := []*Node{root}

	for i := 0; len(currentLevel) > 0; i++ {
		result = append(result, []int{})
		nextLevel := []*Node{}
		for j := 0; j < len(currentLevel); j++ {
			node := currentLevel[j]
			result[i] = append(result[i], node.Val)
			if node.Children != nil {
				nextLevel = append(nextLevel, node.Children...)
			}
		}
		currentLevel = nextLevel
	}
	return result
}
