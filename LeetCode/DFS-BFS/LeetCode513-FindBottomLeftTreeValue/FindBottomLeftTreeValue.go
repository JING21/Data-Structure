package LeetCode513_FindBottomLeftTreeValue

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return root.Val
	}

	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		res = append(res, queue[0].Val)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res[len(res)-1]
}
