package main

func ReverserList(head *Node) *Node {
	    if head.Next == nil{
	    	return head
		}
		LastNode := ReverserList(head.Next)
		head.Next.Next = head
		head.Next =nil
		return LastNode
}


