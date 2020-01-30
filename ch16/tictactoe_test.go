package ch16

import "testing"

import "github.com/magiconair/properties/assert"

func TestTicTacToe(t *testing.T) {
	arr := [][]int{{0, 0, 1, 0}, {1, 1, 0, 1}, {1, 1, 0, 1}, {0, 0, 0, 0}}
	result := hasWonBruteForce(arr)

	assert.Equal(t, result, false)

	arr = [][]int{{0, 0, 1, 0}, {1, 1, 0, 1}, {1, 1, 1, 1}, {0, 0, 0, 0}}
	result = hasWonBruteForce(arr)

	assert.Equal(t, result, true)
}
