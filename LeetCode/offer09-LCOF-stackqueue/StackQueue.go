package offer09_LCOF_stackqueue

type CQueue struct {
	in 		[]int
	out 	[]int
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.in = append(this.in, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0{
		if len(this.in) == 0{
			return -1
		}
		for i := len(this.in); i>0 ; i--{
			this.out = append(this.out, this.in[i-1])
			this.in = this.in[:len(this.in)-1]
		}
		num := this.out[len(this.out)-1]
		this.out = this.out[:len(this.out)-1]
		return  num
	}
	num := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return  num
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */