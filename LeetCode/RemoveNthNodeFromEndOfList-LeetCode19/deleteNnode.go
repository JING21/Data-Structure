package main

type ListNode struct {
     Val int
     Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0,nil}
	dummy.Next = head
	length := 0
	first := head
	for(first != nil){
		length++
		first = first.Next
	}
	length = length-n
	first = dummy
	for (length >0){
		length--
		first = first.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}