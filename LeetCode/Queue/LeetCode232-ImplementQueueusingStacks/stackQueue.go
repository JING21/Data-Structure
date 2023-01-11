package LeetCode232_ImplementQueueusingStacks

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
	inlen, outlen := len(this.stackIn), len(this.stackOut)
	if outlen == 0 {
		if inlen == 0 {
			return -1
		}
		for i := inlen - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}
		outlen = len(this.stackOut)
	}

	val := this.stackOut[outlen-1]
	this.stackOut = this.stackOut[:outlen-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
		return -1
	}

	this.stackOut = append(this.stackOut, val)
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
