package ch04

import "testing"

import "github.com/magiconair/properties/assert"

func TestCommonAncestor(t *testing.T) {

	arr := []int{8, 7, 1, 3, 9, 2, 4, 6, 5, 10, 11}

	root := createNode(arr[0])

	for _, item := range arr[1:] {
		root.insert(item)
	}

	var n1, n2, common *Node

	n1 = root.getNode(10)
	n2 = root.getNode(11)

	common = root.commonAncestor(n1, n2)
	assert.Equal(t, common.data, 10)

	n1 = root.getNode(2)
	n2 = root.getNode(4)

	common = root.commonAncestor(n1, n2)
	assert.Equal(t, common.data, 3)

}
