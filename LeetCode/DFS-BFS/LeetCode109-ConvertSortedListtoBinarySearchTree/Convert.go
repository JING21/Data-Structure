package LeetCode109_ConvertSortedListtoBinarySearchTree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func getMid(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}

	mid := getMid(left, right)
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)

	return root
}