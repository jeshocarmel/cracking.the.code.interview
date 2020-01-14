package ch02

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIntersection(t *testing.T) {
	list1 := createLinkedList([]int{1, 3, 5, 8, 17, 7, 3, 5, 2, 4})
	list2 := createLinkedList([]int{11, 12, 13})

	head1 := list1

	for i := 0; i < 5; i++ {
		head1 = head1.next
	}

	tail2 := list2
	for tail2.next != nil {
		tail2 = tail2.next
	}

	tail2.next = head1

	n, exists := intersection(list1, list2)

	assert.Equal(t, n, head1)
	assert.Equal(t, true, exists)

	list3 := createLinkedList([]int{9, 2, 1, 3})
	n, exists = intersection(list1, list3)

	assert.Equal(t, false, exists)

}
