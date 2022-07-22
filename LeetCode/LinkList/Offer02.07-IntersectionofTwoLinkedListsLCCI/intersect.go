package Offer02_07_IntersectionofTwoLinkedListsLCCI

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0

	for curA != nil {
		curA = curA.Next
		lenA++
	}

	for curB != nil {
		curB = curB.Next
		lenB++
	}

	var step int
	var fast, slow *ListNode

	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}

	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
