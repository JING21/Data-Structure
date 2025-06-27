package LeetCode303_RangeSumQuery

//维护一个sum数组
//其中s[0] = 0
//s[1] = a[0] + a[1]
//s[i] = a[0] + a[1] + .... + a[i-1]
//那么a[left]到a[right]的和可以转换为s[right+1]与s[left]的差

type NumArray struct {
	numArray []int
}

func Constructor(nums []int) NumArray {
	s := NumArray{make([]int, len(nums)+1)}
	for k, v := range nums {
		s.numArray[k+1] = s.numArray[k] + v
	}

	return s
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.numArray[right+1] - this.numArray[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
