package datastructures

import "testing"

import "github.com/magiconair/properties/assert"

func TestBinarySearchTree(t *testing.T) {

	arr := []int{7, 1, 4, 9, 3, 12, 2, 20}

	root := createTreeNode(arr[0])

	for _, item := range arr[1:] {
		root.insert(item)
	}

	var found bool

	// dfs searches
	_, found = root.dfs(3)
	assert.Equal(t, found, true)

	_, found = root.dfs(20)
	assert.Equal(t, found, true)

	_, found = root.dfs(21)
	assert.Equal(t, found, false)

	//bfs searches

	_, found = root.bfs(3)
	assert.Equal(t, found, true)

	_, found = root.bfs(20)
	assert.Equal(t, found, true)

	_, found = root.bfs(21)
	assert.Equal(t, found, false)
}
