package ch03

import "time"

const (
	dog = "dog"
	cat = "cat"
)

type Node struct {
	name      string
	timestamp int
	next      *Node
}

type AnimalShelter struct {
	catlist *Node
	doglist *Node
}

func (as *AnimalShelter) enqueue(name, category string) {

	newnode := &Node{name: name, timestamp: time.Now().Nanosecond()}

	if category == cat {
		if as.catlist == nil {
			as.catlist = newnode
		} else {
			n := as.catlist
			for n.next != nil {
				n = n.next
			}
			n.next = newnode
		}

	} else {
		if as.doglist == nil {
			as.doglist = newnode
		} else {
			n := as.doglist
			for n.next != nil {
				n = n.next
			}
			n.next = newnode
		}
	}
}

func (as *AnimalShelter) dequeueDog() *Node {

	if as.doglist == nil {
		return nil
	}

	dequeued := as.doglist
	as.doglist = as.doglist.next

	return dequeued

}

func (as *AnimalShelter) dequeueCat() *Node {

	if as.catlist == nil {
		return nil
	}

	dequeued := as.catlist
	as.catlist = as.catlist.next

	return dequeued
}

func (as *AnimalShelter) dequeueAny() *Node {

	firstcat := as.catlist
	firstdog := as.doglist

	if firstcat == nil {
		return firstdog
	}

	if firstdog == nil {
		return firstcat
	}

	if firstcat.timestamp < firstdog.timestamp {
		return as.dequeueCat()
	}

	return as.dequeueDog()

}
