package main

type Object interface {}

type Node struct {
	Data Object //定义数据
	Next *Node       //定义地址（指向下一个元素）
}

type LinkList struct {
	headNode *Node  //头结点
}:

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