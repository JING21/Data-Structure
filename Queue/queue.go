package Queue

type Queue struct{
	queue []int
	head  int
	tail  int
	length int
}


func (q *Queue) enqueue (val int) bool{
	if (q.tail == q.length ){
		//尾为n,头为0，表示队列已满
		if q.head == 0 {
			return false
		}
		for i := q.head; i < q.tail; i++{
			q.queue[i-q.head] = q.queue[i]
		}
		q.tail -= q.head
		q.head = 0
	}

	q.queue[q.tail] = val
	q.tail++
	q.length++
	return true
}


func (q *Queue) dequeue () *int{
	if q.head == q.tail{
		return nil
	}
	orgHeadData := q.queue[q.head]
	q.head = q.head+1
	q.length--
	
	return &orgHeadData
}

