package MergeKLinkList_LeetCode25

type ListNode struct {
	Val int
	Next *ListNode
}

// 分治
func mergeKLists(lists []*ListNode) *ListNode {
	ln := len(lists)
	if ln == 0 {
		return nil
	}
	if ln == 1 {
		return lists[0]
	}
	// 分而治之
	return merge2list(mergeKLists(lists[:ln/2]), mergeKLists(lists[ln/2:]))
}

// 合并两个有序链表
func merge2list(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	ans := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ans.Next = list1
			list1 = list1.Next
		} else {
			ans.Next = list2
			list2 = list2.Next
		}
		ans = ans.Next
	}
	for list1 != nil {
		ans.Next = list1
		list1 = list1.Next
		ans = ans.Next
	}
	for list2 != nil {
		ans.Next = list2
		list2 = list2.Next
		ans = ans.Next
	}
	return head.Next
}
