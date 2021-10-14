package offer25_mergeTwoSortList

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}

	if l2 ==nil{
		return l1
	}

	dum :=  &ListNode{-1,nil}
	cur := dum

	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			n1 := l1
			cur.Next = n1
			cur = cur.Next
			l1 = l1.Next
		}else {
			n2 := l2
			cur.Next = n2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 == nil{
		cur.Next = l2
	}

	if l2 == nil{
		cur.Next = l1
	}

	return dum.Next
}