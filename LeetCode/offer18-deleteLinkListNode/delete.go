package offer18_deleteLinkListNode

type ListNode struct {
	Val int
	Next *ListNode
}


func deleteNode(head *ListNode, val int) *ListNode{
	var dummy *ListNode
	dummy = head

	if dummy.Val == val{
		head = dummy.Next
		return head
	}else {
		for dummy.Next != nil{
			if dummy.Next.Val == val{
				dummy.Next = dummy.Next.Next
			}else {
				dummy = dummy.Next
			}
		}
	}
	return head
}


func deleteNode2(head *ListNode, val int) *ListNode{
	if head == nil{
		return nil
	}

	if head.Val == val {
		head = head.Next
		return head
	}

	pre, cur := head, head

	for cur != nil && cur.Val != val  {
		pre, cur = cur, cur.Next
	}

	if cur != nil{
		pre.Next = cur.Next
	}

	return  head
}
