package LeetCode1522_DiameterofN_AryTree

type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {
	result := 0
	var dfs func(node *Node) int
	dfs = func(node *Node) int {
		if node == nil {
			return 0
		}
		dep1, dep2 := 0, 0
		for _, n := range node.Children {
			height := dfs(n)
			if height >= dep1 {
				dep2 = dep1
				dep1 = height
			} else if height > dep2 {
				dep2 = height
			}
		}
		result = max(result, dep1+dep2)
		return max(dep1, dep2) + 1
	}
	dfs(root)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
