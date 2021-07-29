package IntersectionOfTwoLinkedList_LeetCode_interview02_07


//不存在环的情况
type ListNode struct {
	  Val int
	  Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA := 0
	lenB := 0
	for curA !=nil{
		curA = curA.Next
		lenA++
	}
	for curB !=nil{
		curB = curB.Next
		lenB++
	}
	step := 0
	var fast, slow *ListNode
	if lenA > lenB {
		fast, slow = headA, headB
		step = lenA-lenB
	}else {
		fast, slow = headB,headA
		step = lenB-lenA
	}
	for i := 0; i < step; i++{
		fast = fast.Next
	}
	for fast != slow{
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}


//存在环的情况
//1.两链表都不存在环，并且不相交
//2.两链表都不存在环，相交
//3.两链表都存在环，但是相交
//4.两链表一个存在环，一个无环，必定不相交
//5.两链表都存在环，但是相交在环外
//6.两链表都存在环，相交在环上



func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	bool resultA := CheckLinkListRing_LeetCode141.CheckRing()
}