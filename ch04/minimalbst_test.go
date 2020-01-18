package ch04

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMinimalBST(t *testing.T) {

	arr := []int{1, 2, 4, 5, 6, 7, 9, 11, 13, 17, 18, 19, 55, 78, 99}
	root := createMinimalBST(arr)
	height := root.maxdepth()

	assert.Equal(t, height, 4)
}
