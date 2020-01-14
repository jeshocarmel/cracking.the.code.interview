package ch02

import "testing"

import "github.com/magiconair/properties/assert"

func TestKthtoLast(t *testing.T) {

	head := createLinkedList([]int{2, 4, 1, 3, 6, 8, 1, 3, 5, 7})
	ans := kthToLast(head, 4)

	assert.Equal(t, ans, 8)

	ans = kthToLast(head, 1)
	assert.Equal(t, ans, 5)

	ans = kthToLast(head, 9)
	assert.Equal(t, ans, 2)

}
