package MinStack_LeetCode155

import "math"

type MinStack struct {
	stack []int
	minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack,min(top, val))
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int{
	if x < y{
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */


type MinStackO1 struct {
	stack []int
	min   int
}


/** initialize your data structure here. */
func Constructor2() MinStackO1 {
	return MinStackO1{
		stack: []int{},
		min:   0,
	}
}


func (this *MinStackO1) Push2(val int)  {
	if len(this.stack) == 0 {
		this.min = val
		this.stack = append(this.stack, val-this.min)
	}else if val - this.min < 0 {
		this.stack = append(this.stack, val-this.min)
		this.min = val
	}else {
		this.stack = append(this.stack, val-this.min)
	}
}


func (this *MinStackO1) Pop2()  {
	if this.stack[len(this.stack)-1] < 0 {
		this.min = this.min - this.stack[len(this.stack)-1]

	}
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStackO1) Top2() int {
	if this.stack[len(this.stack)-1] < 0 {
		return this.min
	}
	return this.min+this.stack[len(this.stack)-1]
}


func (this *MinStackO1) GetMin2() int {
	return this.min
}

