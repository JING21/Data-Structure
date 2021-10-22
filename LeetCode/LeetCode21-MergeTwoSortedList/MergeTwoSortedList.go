package LeetCode21_MergeTwoSortedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	     Val int
	    Next *ListNode
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//如果l1是nil,直接返回l2，因为链表本身是有序的
	if l1 == nil{
		return l2
	}
	//如果l2是nil,直接返回l2，因为链表本身是有序的
	if l2  == nil{
		return l1
	}
	pre := &ListNode{-1,nil}
	prehead := pre
	for (l1 != nil && l2 != nil){
		if (l1.Val < l2.Val){
			n1 := l1
			pre.Next = n1
			pre = pre.Next
			l1 = l1.Next
		} else {
			n2 := l2
			pre.Next = n2
			pre = pre.Next
			l2 = l2.Next
		}

	}

	if (l1 == nil) {
		pre.Next = l2
	} else if (l2 == nil){
		pre.Next =l1
	}
	return prehead.Next
}
