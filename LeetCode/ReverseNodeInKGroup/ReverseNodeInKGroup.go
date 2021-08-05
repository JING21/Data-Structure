package ReverseNodeInKGroup

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	dummy := &ListNode{0,head}
	tem := dummy
	len := 0
	for head != nil {
		head = head.Next
		len++
	}
	 for i := 0; i+k <= len; i+=k {
	 		cur := tem.Next

	 }
}