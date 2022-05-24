package LeetCode993_CousinOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	node  *TreeNode
	depth int
}

//BFS
func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	update := func(node, parent *TreeNode, depth int) {
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
	}

	nextLevel := []pair{{root, 0}}
	update(root, nil, 0)
	for len(nextLevel) > 0 && (!xFound || !yFound) {
		node, depth := nextLevel[0].node, nextLevel[0].depth
		nextLevel = nextLevel[1:]
		if node.Left != nil {
			nextLevel = append(nextLevel, pair{node.Left, depth + 1})
			update(node.Left, node, depth+1)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, pair{node.Right, depth + 1})
			update(node.Right, node, depth+1)
		}

	}
	return xDepth == yDepth && xParent != yParent
}

func isCousins2(root *TreeNode, x, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	// 用来判断是否遍历到 x 或 y 的辅助函数
	update := func(node, parent *TreeNode, depth int) {
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
	}

	type pair struct {
		node  *TreeNode
		depth int
	}
	q := []pair{{root, 0}}
	update(root, nil, 0)
	for len(q) > 0 && (!xFound || !yFound) {
		node, depth := q[0].node, q[0].depth
		q = q[1:]
		if node.Left != nil {
			q = append(q, pair{node.Left, depth + 1})
			update(node.Left, node, depth+1)
		}
		if node.Right != nil {
			q = append(q, pair{node.Right, depth + 1})
			update(node.Right, node, depth+1)
		}
	}

	return xDepth == yDepth && xParent != yParent
}

//DFS
func DFS(root *TreeNode, x, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xParent, xDepth, xFound = node, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}

		if xFound && yFound {
			return
		}

		dfs(node.Left, node, depth+1)

		if xFound && yFound {
			return
		}

		dfs(node.Right, node, depth+1)

	}
	dfs(root, nil, 0)

	return xDepth == yDepth && xParent != yParent
}
