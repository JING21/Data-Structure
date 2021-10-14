package offer52_IntersectionOfTwoLinkList


type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB

	if headA == nil || headB == nil{
		return nil
	}

	for nodeA != nodeB{
		if nodeA == nil{
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}
		if nodeB == nil{
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nodeA
}