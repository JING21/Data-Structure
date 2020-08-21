package main

import (
	"fmt"
	"github.com/juju/errors"
)

type Stack struct {
	MaxNum int   //栈所包含的数量个数
	Top    int   //栈顶数据
	arr    []int //模拟栈
}

func (s *Stack) Push(val int) error{
	// 先判断栈是否满了
	if s.Top == s.MaxNum-1{
		fmt.Println("The stack is full")
		return errors.New("The stack is full")
	}
	s.Top++
	s.arr[s.Top] = val
	return nil
}

func (s *Stack) Push2(val int) (err error) {
	if s.Top == s.MaxNum-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	s.Top++
	// 放入数据
	s.arr[s.Top] = val
	return

}

func (s *Stack) Pop() error{
	if s.Top == -1 {
		fmt.Println("The stack is empty")
		return errors.New("The stack is empty")
	}
	val := s.arr[s.Top]
	s.Top--
	fmt.Println("")
	fmt.Println("***************")
	fmt.Printf("Stack pop %v",val)
	fmt.Println("")
	fmt.Println("***************")
	return nil
}

func (s *Stack) List() error{
	if s.Top == -1{
		fmt.Println("The stack is empty")
		return errors.New("The stack is empty")
	}
	fmt.Println("The stack is")
	for i := s.Top; i >= 0; i--{
		fmt.Printf("Stack[%d]=%d",i, s.arr[i])
	}
	return nil
}


func main() {
	slice := make([]int,5)
	stack := &Stack{
		MaxNum: 5,
		Top:    -1,
		arr: slice,
	}
	// 入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	// 显示
	stack.List()
	stack.Pop()
	// 显示
	stack.List()
}
