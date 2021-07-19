package main

func CheckRing(head *Node) *Node {
	m := make(map[*Node]Object)
	for head != nil {
		_, ok := m[head]
		if ok == false {
			m[head] = 1
			head = head.Next
		}else {
			return head
		}
	}
	return nil
}


func CheckRingTwo(head *Node) bool{
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

