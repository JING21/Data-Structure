package Queue

import "fmt"

type ListNode struct {
	val 	interface{}
	next 	*ListNode
}

type ListNodeQueue struct {
	head *ListNode
	tail *ListNode
	length int
}

func NewLinkedListQueue() *ListNodeQueue{
	return &ListNodeQueue{nil, nil, 0}
}


func (q *ListNodeQueue) Enqueue (v interface{}){
	node := &ListNode{v,nil}
	if q.tail == nil{
		q.tail = node
		q.head = node
	}else {
		q.tail.next = node
		q.tail = node
	}
	q.length ++
}


func (q *ListNodeQueue) DeQueue () interface{}{
	if q.head == nil{
		return nil
	}
	v := q.head.val
	q.head = q.head.next
	q.length++
	return v
}

func (q *ListNodeQueue) Print() string{
	if q.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
