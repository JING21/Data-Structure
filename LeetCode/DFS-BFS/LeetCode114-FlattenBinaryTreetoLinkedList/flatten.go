package LeetCode114_FlattenBinaryTreetoLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func flatten(root *TreeNode) {
//	list := preorderTraversal(root)
//	for i := 1; i < len(list); i++ {
//		prev, curr := list[i-1], list[i]
//		prev.Left, prev.Right = nil, curr
//	}
//}
//
//func preorderTraversal(root *TreeNode) []*TreeNode {
//	list := []*TreeNode{}
//	if root != nil {
//		list = append(list, root)
//		list = append(list, preorderTraversal(root.Left)...)
//		list = append(list, preorderTraversal(root.Right)...)
//	}
//	return list
//}

func flatten(root *TreeNode) {
	result := []*TreeNode{}

	if root == nil {
		return
	}

	var dfs func(*TreeNode) []*TreeNode

	dfs = func(node *TreeNode) []*TreeNode {
		if node == nil {
			return result
		}
		result = append(result, node)
		dfs(node.Left)
		dfs(node.Right)

		return result
	}
	rs := dfs(root)

	for i := 1; i < len(rs); i++ {
		pre, curr := rs[i-1], rs[i]
		pre.Left, pre.Right = nil, curr
	}
}

func flattenMorris(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
