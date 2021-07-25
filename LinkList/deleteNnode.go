package main


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

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0,nil}
	dummy.Next = head
	fast, slow := head, dummy
	for i := 0; i <n ; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}