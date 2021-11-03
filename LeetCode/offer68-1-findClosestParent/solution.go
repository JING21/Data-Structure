package offer68_1_findClosestParent

type TreeNode struct {
	 Val   int
  	 Left  *TreeNode
  	 Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor := root

	for{
		if p.Val < ancestor.Val && q.Val < ancestor.Val{
			ancestor = ancestor.Left
		}else if p.Val > ancestor.Val && q.Val > ancestor.Val{
			ancestor = ancestor.Right
		}else {
			return ancestor
		}
	}
}

