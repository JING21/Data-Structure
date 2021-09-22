package Queue

type CycleQueue struct {
	queue 	   []interface{}
	capacity   int
	head       int
	tail	   int
}


func NewCycleQueue(n int) *CycleQueue{
	if n == 0 {
		return nil
	}
	return &CycleQueue{make([]interface{}, n), n ,0,0}
}


func (q *CycleQueue) IsEmpty() bool{
	if q.head == q.tail{
		return true
	}
	return false
}

//如果tail+1 % capacity == head
func (q *CycleQueue) IsFull() bool{
	if q.head == (q.tail+1)%q.capacity{
		return true
	}
	return false
}


func (q *CycleQueue) EnQueue(val  interface{}) bool{
	if q.IsFull() == true{
		return false
	}
	q.queue[q.tail] = val
	q.tail = (q.tail+1) % q.capacity
	return true
}

func (q *CycleQueue) DeQueue() interface{}{
	if q.IsEmpty() == true{
		return false
	}
	val := q.queue[q.head]
	q.head = (q.head+1) % q.capacity
	return val
}

