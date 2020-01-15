package datastructures

import (
	"testing"
)

func TestStack(t *testing.T) {

	mystack := NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	mystack.Push(5)

	mystack.Pop()
	mystack.Pop()

	mystack.Push(6)

	mystack.Print()

}
