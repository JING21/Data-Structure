package LeetCode559_MaximumDepthOfN_aryTree

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	maxChildDepth := 0

	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}

func maxDepth2(root *Node) (ans int) {
	if root == nil {
		return 0
	}
	nextLevel := []*Node{root}

	for len(nextLevel) > 0 {
		currentLevel := nextLevel
		nextLevel = nil
		for _, node := range currentLevel {
			nextLevel = append(nextLevel, node.Children...)
		}
		ans++
	}
	return
}
