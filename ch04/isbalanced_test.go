package ch04

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIsBalanced(t *testing.T) {

	var arr []int
	var root *Node

	//1st scenario
	arr = []int{50, 17, 76, 9, 23, 54, 14, 19, 72, 12, 67}
	root = createNode(arr[0])

	for _, item := range arr[1:] {
		root.insert(item)
	}

	isBalanced := root.isBalanced()
	assert.Equal(t, isBalanced, false)

	//2nd scenario
	arr = []int{50, 17, 72, 12, 23, 54, 76, 9, 14, 19, 67}
	root = createNode(arr[0])

	for _, item := range arr[1:] {
		root.insert(item)
	}

	isBalanced = root.isBalanced()
	assert.Equal(t, isBalanced, true)
}
