package LeetCode257_BinaryTreePaths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var paths []string
//
//func binaryTreePaths(root *TreeNode) []string {
//	paths = []string{}
//	constructPaths(root, "")
//	return paths
//}
//
//func constructPaths(root *TreeNode, path string) {
//	if root != nil {
//		pathSB := path
//		pathSB += strconv.Itoa(root.Val)
//		if root.Left == nil && root.Right == nil {
//			paths = append(paths, pathSB)
//		} else {
//			pathSB += "->"
//			constructPaths(root.Left, pathSB)
//			constructPaths(root.Right, pathSB)
//		}
//	}
//}

var result []string

func binaryTreePaths(root *TreeNode) []string {
	result = []string{}
	DFS(root, "")
	return result
}

func DFS(root *TreeNode, path string) {
	if root != nil {
		pathSub := path
		pathSub += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			result = append(result, pathSub)
		} else {
			pathSub += "->"
			DFS(root.Left, pathSub)
			DFS(root.Right, pathSub)
		}
	}
}
