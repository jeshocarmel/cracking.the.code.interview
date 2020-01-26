package ch04

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPathSum(t *testing.T) {
	root := createNode(10)

	insert(root, 5)
	insert(root, -3)
	insert(root, 2)
	insert(root, 6)
	insert(root, 11)
	insert(root, 1)
	insert(root, 8)

	ans := getNoOfPaths(root, 2)
	assert.Equal(t, ans, 2)
}
