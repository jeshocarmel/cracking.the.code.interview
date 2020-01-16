package datastructures

import "fmt"

//Queue represents the queue datastructure
type Queue struct {
	first *Item
	last  *Item
}

//NewQueue returns a new queue
func NewQueue() *Queue {
	q := new(Queue)
	return q
}

//Add adds to the queue
func (q *Queue) Add(i interface{}) {

	newItem := &Item{data: i}

	if q.first == nil {
		q.first = newItem
		q.last = newItem
	} else {
		q.last.next = newItem
		q.last = newItem
	}
}

//Remove removes the first item from the queue
func (q *Queue) Remove() interface{} {

	if q.IsEmpty() {
		return nil
	}
	deleted := q.first
	q.first = q.first.next
	return deleted.data
}

//Peek returns the last stored item in the queue
func (q *Queue) Peek() interface{} {

	if q.last == nil {
		return nil
	}

	return q.last.data
}

//IsEmpty returns if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

//Print prints the queue
func (q *Queue) Print() {

	fmt.Print("\nqueue ==> ")
	for n := q.first; n != nil; n = n.next {

		fmt.Print(n.data)
		fmt.Print(" ")

	}
	fmt.Print("\n")
}
