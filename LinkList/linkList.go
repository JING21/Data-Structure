package main

import "fmt"

type Object interface {}

type Node struct {
	Data Object //定义数据
	Next *Node       //定义地址（指向下一个元素）
}

type LinkList struct {
	headNode *Node  //头结点
}

func (L *LinkList) IsEmpty() bool {
	if L.headNode == nil {
		return true
	} else {
		return false
	}
}


func (L *LinkList) Length() int {
	cur := L.headNode
	count := 0

	for cur != nil{
		count++
		cur = cur.Next
	}
	return count
}

//头部添加元素
func (L *LinkList)AddHead(data Object) {
	node := &Node{Data: data}
	node.Next = L.headNode
	L.headNode = node

}

//尾部添加元素
func (L *LinkList) AddLast(data Object) {
	node := &Node{Data: data}
	//判断是否为空
	if L.IsEmpty() {
		L.headNode = node
	} else {
		cur := L.headNode
		//确认为最后一个节点
		for cur.Next != nil{
			cur = cur.Next
		}
		cur.Next = node
	}
}

//添加元素在链表的具体位置
func (L *LinkList) Insert(index int, data Object) {
	if index < 0 {
		L.AddHead(data)
	} else if index > L.Length() {
		L.AddLast(data)
	} else {
		preHead := L.headNode
		count := 0
		for count < (index-1){
			preHead = preHead.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = preHead.Next
		preHead.Next = node
	}

}

//删除值为data的元素
func (L *LinkList) Delete(data Object){
	preHead := L.headNode
	if preHead.Data == data{
		L.headNode = preHead.Next
	} else {
		for preHead.Next != nil{
			if preHead.Next.Data == data {
				preHead.Next = preHead.Next.Next
			} else {
				preHead = preHead.Next
			}
		}
	}
}

//删除下标为index的具体元素
func (L *LinkList) RemoveIndex(index int){
	preHead := L.headNode
	if index == 0 {
		L.headNode =preHead.Next
	} else if index > L.Length() {
		fmt.Println("It is over the length of the linklist")
	} else {
		count := 0
		for count != (index-1) && preHead.Next != nil{
			count++
			preHead = preHead.Next
		}
		preHead.Next = preHead.Next.Next
	}
}

//是否包含该值的元素
func (L *LinkList) Contain(data Object) bool{
	cur := L.headNode
	for cur != nil {
		if cur.Data == data{
			return true
		}
		cur = cur.Next
	}
	return false
}

//遍历链表
func (L *LinkList) ShowList(){
	if !L.IsEmpty(){
		cur := L.headNode
		for {
			fmt.Println("\t%v",cur.Data)
			if cur.Next != nil{
				cur = cur.Next
			}else {
				break
			}
		}
	}
}