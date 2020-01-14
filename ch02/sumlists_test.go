package ch02

import "testing"

import "github.com/magiconair/properties/assert"

func TestSumListsReverse(t *testing.T) {

	list1 := createLinkedList([]int{7, 1, 6})
	list2 := createLinkedList([]int{5, 9, 2})

	ans := sumListsReverse(list1, list2)

	assert.Equal(t, ans, 912)

	list1 = createLinkedList([]int{7, 1, 6, 9, 1})
	list2 = createLinkedList([]int{5, 9, 2, 5})

	ans = sumListsReverse(list1, list2)
	assert.Equal(t, ans, 24912)

}

func TestSumListsNormal(t *testing.T) {

	list1 := createLinkedList([]int{6, 1, 7})
	list2 := createLinkedList([]int{2, 9, 5})

	ans := sumListsNormal(list1, list2)

	assert.Equal(t, ans, 912)

	list1 = createLinkedList([]int{1, 9, 6, 1, 7})
	list2 = createLinkedList([]int{5, 2, 9, 5})

	ans = sumListsNormal(list1, list2)
	assert.Equal(t, ans, 24912)

}
