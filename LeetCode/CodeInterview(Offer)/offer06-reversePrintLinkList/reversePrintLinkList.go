package offer06_reversePrintLinkList

type ListNode struct {
		Val int
  		Next *ListNode
}

func reversePrint(head *ListNode) []int {
	result := reverseLinkList(head)
	var finalResult []int
	for result.Next != nil{
		finalResult = append(finalResult, result.Val)
		result = result.Next
	}
	return finalResult
}


func reverseLinkList(head *ListNode) *ListNode{
	if head.Next == nil{
		return head
	}
	LastNode := reverseLinkList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return LastNode
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint2(head *ListNode) []int {
	var result []int
	if head == nil {
		return nil
	}
	for head != nil{
		result = append(result, head.Val)
		head = head.Next
	}
	length := len(result)
	for i := 0; i < length/2; i++ {
		temp := result[length-1-i]
		result[length-1-i] = result[i]
		result[i] = temp
	}
	return result
}