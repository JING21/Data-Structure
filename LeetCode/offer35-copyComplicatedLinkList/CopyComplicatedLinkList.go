package offer35_copyComplicatedLinkList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node{
	if head == nil{
		return nil
	}

	for origin := head; origin != nil;{
		cp := &Node{origin.Val, origin.Next, nil}
		origin.Next = cp

		origin = cp.Next
	}

	for origin := head; origin != nil; {
		if origin.Random != nil{
			origin.Next.Random = origin.Random.Next
		}
		origin = origin.Next.Next
	}

	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}

	return headNew

}