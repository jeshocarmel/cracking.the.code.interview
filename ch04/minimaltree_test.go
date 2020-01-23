package ch04

import "testing"

import "github.com/stretchr/testify/assert"

func TestMinimalTree(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	root := divideAndConquer(arr, 0, len(arr), nil)

	inOrderStr := inOrder(root)
	assert.Equal(t, inOrderStr, "1 2 3 4 5 6 7 8 9 ")
	assert.Equal(t, 5, root.data)

}
