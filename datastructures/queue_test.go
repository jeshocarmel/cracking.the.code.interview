package datastructures

import (
	"testing"
)

func TestQueue(t *testing.T) {

	myqueue := NewQueue()

	myqueue.Add(1)
	myqueue.Add(2)
	myqueue.Add(3)
	myqueue.Add(4)
	myqueue.Remove()
	myqueue.Remove()
	myqueue.Add(5)
	myqueue.Remove()
	myqueue.Remove()
	myqueue.Add(6)

	myqueue.Print()

}
