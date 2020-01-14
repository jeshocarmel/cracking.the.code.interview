package ch02

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	node := createLinkedList([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})
	ans := isPalindromeUsingStack(node)
	assert.Equal(t, ans, true)

	node = createLinkedList([]int{1, 2, 3, 4, 4, 3, 2, 1})
	ans = isPalindromeUsingStack(node)
	assert.Equal(t, ans, true)

	node = createLinkedList([]int{1, 2, 3, 4, 9, 5, 6, 9, 4, 3, 2, 1})
	ans = isPalindromeUsingStack(node)
	assert.Equal(t, ans, false)

}

func TestIsPalindromeRecursive(t *testing.T) {
	node := createLinkedList([]int{1, 2, 3, 4, 9, 8, 4, 3, 2, 1})
	isPalindromeRecursive(node)

}
