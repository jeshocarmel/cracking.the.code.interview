package ch04

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

/*
	instead of using linked-list this example uses maps.
*/
func TestListOfDepths(t *testing.T) {

	arr := []int{8, 1, 3, 9, 12, 14, 5, 16, 29}

	root := createNode(arr[0])

	for _, item := range arr[1:] {
		root.insert(item)
	}

	dfsLevelMap := make(map[int][]int)
	root.dfsTraversal(dfsLevelMap, 0)

	bfsLevelMap := make(map[int][]int)
	root.bfsTraversal(bfsLevelMap)

	assert.Equal(t, dfsLevelMap, bfsLevelMap)
}
