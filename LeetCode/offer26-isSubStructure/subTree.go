package offer26_isSubStructure
//
type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}
//
//func isSubStructure(A *TreeNode, B *TreeNode) bool {
//
//	// 根据题目判断条件，初始有空节点返回false
//	if A == nil || B == nil{
//		return false
//	}
//
//	queue, root := []*TreeNode{A}, A
//
//	for len(queue) > 0{
//
//		root = queue[0]
//		queue = queue[1:]
//
//		if root.Val == B.Val && helper(root, B){
//			return true
//		}
//
//		// 还没遍历完，而且没有找到匹配的子树，继续遍历A子树
//		if root.Left != nil {
//			queue = append(queue, root.Left)
//		}
//		if root.Right != nil {
//			queue = append(queue, root.Right)
//		}
//	}
//
//	return false
//
//}
//
//func helper (A *TreeNode, B *TreeNode) bool {
//
//	if B == nil {
//		return true
//	}
//
//	if B == nil || A == nil {
//		return false
//	}
//
//	if A.Val == B.Val {
//		return helper(A.Left, B.Left) &&  helper(A.Right, B.Right)  //递归
//	}else{
//		return false
//	}
//}


func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil{
		return false
	}
	queue := []*TreeNode{A}

	for len(queue) >0 {
		currentNode := queue[0]
		if currentNode.Val == B.Val{
			if containSubTree(currentNode, B){
				return true
			}
		}

		queue = queue[1:]

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
	return false
}


func containSubTree(A *TreeNode, B *TreeNode) bool{
	if B == nil{
		return true
	}
	if A == nil{
		return false
	}
	if A.Val == B.Val{
		return containSubTree(A.Left, B.Left) && containSubTree(A.Right, B.Right)
	}else {
		return false
	}
}