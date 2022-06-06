package LeetCode426_ConvertBinarySearchTreetoSortedDoublyLinkedList

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList2(root *Node) *Node {
	if root == nil {
		return nil
	}
	first := &Node{}
	last := first
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		dfs(node.Left)
		last.Right = node
		node.Left = last
		last = last.Right
		dfs(node.Right)
	}
	dfs(root)
	//构造环
	head := first.Right
	head.Left = last
	last.Right = head
	return head

}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	var first, last *Node
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if last != nil {
			last.Right = node
			node.Left = last
		} else {
			first = node
		}
		last = node
		dfs(node.Right)
	}
	dfs(root)
	//构造环
	last.Right = first
	first.Left = last
	return first
}

// 左根右
func treeToDoublyList3(root *Node) *Node {
	// 前驱节点，当前头节点
	var pre, head *Node
	stack := []*Node{}
	for root != nil || len(stack) > 0 {
		// 中序遍历每次都从当前节点的最左孩子开始处理
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈处理当前节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 首次遍历，前驱节点为空，则记录头节点
		if pre == nil {
			head = root
		} else {
			// 非首次遍历，前驱节点与刚出栈的根节点相互指向建立连接
			pre.Right = root
			root.Left = pre
		}
		// 前驱节点记录当前根节点，切换到右孩子进行处理
		pre = root
		root = root.Right
	}
	// 单向循环结束以后，需要将最左孩子与最右孩子相互指向建立连接
	if pre != nil && head != nil {
		pre.Right = head
		head.Left = pre
	}
	return head
}
