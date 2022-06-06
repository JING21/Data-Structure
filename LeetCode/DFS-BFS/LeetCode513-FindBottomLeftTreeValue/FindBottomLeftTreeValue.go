package LeetCode513_FindBottomLeftTreeValue

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	result := [][]int{}

	if root == nil {
		return result[0][0]
	}

	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {
		result = append(result, []int{})
		q := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			result[i] = append(result[i], node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		queue = q

	}
	return result[len(result)-1][0]
}

var max, result int

func findBottomLeftValue2(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	DFS(root, max)
	return result
}
func DFS(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		if depth > max {
			result = root.Val
			max = depth
		}
	}
	if root.Left != nil {
		DFS(root.Left, depth+1)
	}
	if root.Right != nil {
		DFS(root.Right, depth+1)
	}
}
