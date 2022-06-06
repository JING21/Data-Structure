package LeetCode199_BinaryTreeRightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res = []int{}
	DFS(root, 1)
	return res
}

func DFS(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level > len(res) {
		res = append(res, root.Val)
	}

	DFS(root.Right, level+1)
	DFS(root.Left, level+1)
}

func rightSideViewBFS(root *TreeNode) []int {
	curr := []*TreeNode{root}
	result := []int{}

	if root == nil {
		return result
	}

	for len(curr) > 0 {
		size := len(curr)
		result = append(result, curr[size-1].Val)

		for i := 0; i < size; i++ {
			node := curr[i]

			if node.Left != nil {
				curr = append(curr, node.Left)
			}
			if node.Right != nil {
				curr = append(curr, node.Right)
			}
		}
		curr = curr[size:]
	}
	return result
}

var res []int
