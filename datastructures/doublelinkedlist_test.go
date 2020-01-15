package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedList(t *testing.T) {

	nodeA := CreateNode(10, 100)
	nodeB := CreateNode(11, 110)
	nodeC := CreateNode(12, 120)
	nodeD := CreateNode(13, 130)
	nodeE := CreateNode(14, 140)
	nodeF := CreateNode(1, 10)
	nodeG := CreateNode(100, 1000)

	var mydll DoubleLinkedList

	mydll.Insert(nodeA)
	mydll.Insert(nodeB)
	mydll.Insert(nodeC)

	mydll.Insertafter(nodeB, nodeD)
	mydll.InsertBefore(nodeD, nodeE)
	mydll.InsertAtBeginning(nodeF)
	mydll.InsertAtEnd(nodeG)

	mydll.Remove(nodeE)

	assert.Equal(t, mydll.GetKeysAsList(), []int{1, 10, 11, 13, 12, 100})

}
