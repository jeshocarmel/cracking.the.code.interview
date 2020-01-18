package ch04

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsBST(t *testing.T) {

	var arr []int
	var root *Node

	// first option - iterative approach

	arr = []int{8, 1, 10}
	root = createNode(arr[0])

	for _, item := range arr[1:] {
		root.insert(item)
	}
	assert.Equal(t, root.isBST(), true)

	//second option - recursive approach

	arr = []int{8, 1, 10, 1}
	root = createNode(arr[0])

	for _, item := range arr[1:] {
		root.insert(item)
	}
	assert.Equal(t, root.isBSTRecursive(nil, nil), true)

}
