package ch16

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDoubleLinkedList(t *testing.T) {

	var mydll DoubleLinkedList
	nodeA := mydll.insert(10)
	nodeB := mydll.insert(9)
	nodeC := mydll.insert(8)

	mydll.insertafter(6, nodeB)
	mydll.insertafter(11, nodeC)
	nodeE := mydll.insertBefore(4, nodeC)
	mydll.insertBefore(1, nodeA)
	mydll.insertBefore(19, nodeE)

	mydll.insertAtBeginning(10)
	mydll.insertAtEnd(20)

	assert.Equal(t, mydll.getAsList(), []int{10, 1, 10, 9, 6, 19, 4, 8, 11, 20})

}
