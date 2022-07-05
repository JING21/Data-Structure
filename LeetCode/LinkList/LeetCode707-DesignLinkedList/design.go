package LeetCode707_DesignLinkedList

type MyLinkedList struct {
	DummyHead *Node
	DummyTail *Node
	Size      int
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func Constructor() MyLinkedList {
	dummyHead := &Node{-1, nil, nil}
	dummyTail := &Node{-1, nil, dummyHead}
	dummyHead.Next = dummyTail
	return MyLinkedList{dummyHead, dummyTail, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	} else {
		p := this.DummyHead
		for i := 0; i <= index; i++ {
			p = p.Next
		}
		return p.Val
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{val, nil, nil}
	newNode.Next = this.DummyHead.Next
	newNode.Pre = this.DummyHead
	this.DummyHead.Next.Pre = newNode
	this.DummyHead.Next = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{val, nil, nil}
	newNode.Next = this.DummyTail
	newNode.Pre = this.DummyTail.Pre
	this.DummyTail.Pre.Next = newNode
	this.DummyTail.Pre = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	} else if index < 0 {
		this.AddAtHead(val)
	} else {
		p := this.DummyHead
		for i := 0; i < index; i++ {
			p = p.Next
		}
		newNode := &Node{val, p.Next, p}
		p.Next.Pre = newNode
		p.Next = newNode
		this.Size++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size || index < 0 {
		return
	} else {
		p := this.DummyHead
		for i := 0; i < index; i++ {
			p = p.Next
		}
		p.Next.Next.Pre = p
		p.Next = p.Next.Next
		this.Size--
	}
}
