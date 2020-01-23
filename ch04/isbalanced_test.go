package ch04

import "testing"

import "github.com/magiconair/properties/assert"

func TestIsBalancedBF(t *testing.T) {

	root := createNode(50)
	arr := []int{25, 75, 37, 63, 13, 87}

	for _, item := range arr {
		insert(root, item)
	}

	var ans bool
	ans = isBalancedBruteForce(root)
	assert.Equal(t, ans, true)

}

func TestIsBalancedOptimized(t *testing.T) {

	root := createNode(50)
	arr := []int{25, 75, 37, 63, 13, 87}

	for _, item := range arr {
		insert(root, item)
	}

	var ans bool
	ans = isBalancedOptimized(root)
	assert.Equal(t, ans, true)

}
