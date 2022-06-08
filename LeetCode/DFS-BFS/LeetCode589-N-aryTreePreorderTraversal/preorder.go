package LeetCode589_N_aryTreePreorderTraversal

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

//func preorder(root *Node) []int {
//	result := []int{}
//	if root == nil {
//		return result
//	}
//
//	var dfs func(node *Node)
//	dfs = func(node *Node) {
//		if node == nil {
//			return
//		}
//		children
//		result = append(result)
//	}
//}

func preorder(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return
}
