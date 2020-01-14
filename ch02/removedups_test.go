package ch02

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRemoveDups(t *testing.T) {

	head := createLinkedList([]int{2, 4, 5, 1, 4, 5, 8, 2})

	ans := removeDups(head).getAsList()
	assert.Equal(t, ans, []int{2, 4, 5, 1, 8})

}

func TestRemoveDupsOptimized(t *testing.T) {

	head := createLinkedList([]int{2, 4, 5, 1, 4, 5, 8, 2})

	ans := removeDupsOptimized(head).getAsList()
	assert.Equal(t, ans, []int{2, 4, 5, 1, 8})

}
