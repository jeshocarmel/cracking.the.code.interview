package ch02

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {

	head := createLinkedList([]int{6, 2, 4, 5, 1})
	middle := head.next.next

	deleteMiddleNode(middle)

	ans := head.getAsList()
	assert.Equal(t, ans, []int{6, 2, 5, 1})
}
