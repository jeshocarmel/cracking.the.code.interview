package datastructures

import "fmt"

//credits - https://github.com/krlv/ctci-go/blob/master/ch03/stack.go

//Item represents the item in a stack
type Item struct {
	data interface{}
	next *Item
}

//Stack represents the datastructure
type Stack struct {
	top *Item
}

//NewStack creates and returns a new stack
func NewStack() *Stack {
	s := new(Stack)
	return s
}

//Push pushes an item into the stack
func (s *Stack) Push(i interface{}) {

	newItem := &Item{data: i}

	if s.top == nil {
		s.top = newItem
	} else {
		newItem.next = s.top
		s.top = newItem
	}

}

//Pop pops an item out of the stack
func (s *Stack) Pop() interface{} {

	if s.IsEmpty() {
		return nil
	}

	deleted := s.top
	s.top = s.top.next
	return deleted.data

}

//Peek returns the top element in the stack
func (s *Stack) Peek() interface{} {
	return s.top.data
}

//IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

//Print prints the stack
func (s *Stack) Print() {

	top := s.top

	fmt.Print("\nstack ==> ")

	for top != nil {
		fmt.Print(top.data)
		fmt.Print(" ")
		top = top.next
	}

	fmt.Print("\n")

}
