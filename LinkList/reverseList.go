package main

func ReverserList(head *Node) *Node {
	    if head.Next == nil{
	    	return head
		}
		LastNode := ReverserList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return LastNode
}

var successor *Node

func ReverseNList(head *Node, n int) *Node {
	if (n == 1) {
		finalNode := head.Next
		return finalNode
	}
	lastNode := ReverseNList(head.Next, n-1)

	head.Next.Next = head
	head.Next = successor
	return lastNode
}