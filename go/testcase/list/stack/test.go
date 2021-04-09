package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	top *ListNode // 栈顶指针
	len int
}

func NewStack() *Stack {
	return &Stack{
		top: nil,
		len: 0,
	}
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Pop() (bool, int) {
	if s.len == 0 {
		return false, 0
	}

	tmp := s.top
	s.top = tmp.Next
	s.len --
	return true, tmp.Val
}

func (s *Stack) Push(val int) {
	tmp := &ListNode{val, nil}
	if s.len == 0 {
		s.top = tmp
	} else {
		tmp.Next = s.top
		s.top = tmp
	}
	s.len ++
}

func (s *Stack) print() {
	node := s.top
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.print()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
