package ch02

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPartition(t *testing.T) {

	head := createLinkedList([]int{3, 5, 8, 5, 10, 2, 1})
	ans := partition(head, 5)

	assert.Equal(t, ans.getAsList(), []int{1, 2, 3, 10, 5, 8, 5})

	ans = partition2(head, 5)
	assert.Equal(t, ans.getAsList(), []int{1, 2, 3, 5, 8, 5, 10})
}
