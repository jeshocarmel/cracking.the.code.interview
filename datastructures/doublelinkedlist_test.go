package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedList(t *testing.T) {

	var mydll DoubleLinkedList
	nodeA := mydll.Insert(10, 100)
	nodeB := mydll.Insert(9, 90)
	nodeC := mydll.Insert(8, 80)

	mydll.Insertafter(6, 60, nodeB)
	mydll.Insertafter(11, 110, nodeC)
	nodeE := mydll.InsertBefore(4, 40, nodeC)
	mydll.InsertBefore(1, 10, nodeA)
	mydll.InsertBefore(19, 190, nodeE)

	mydll.InsertAtBeginning(17, 170)
	mydll.InsertAtEnd(20, 200)

	mydll.Remove(4)

	assert.Equal(t, mydll.GetKeysAsList(), []int{17, 1, 10, 9, 6, 19, 8, 11, 20})

}
