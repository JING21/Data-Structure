package LeetCode589_N_aryTreePreorderTraversal

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, children := range node.Children {
			dfs(children)
		}
	}
	dfs(root)

	return result
}

