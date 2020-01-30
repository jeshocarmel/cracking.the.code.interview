package ch16

import "testing"

import "github.com/magiconair/properties/assert"

func TestTicTacToe(t *testing.T) {
	arr := [][]int{{0, 1, 1}, {0, 0, 1}, {1, 0, 0}}
	result := hasWon(arr, 1, 1)
	assert.Equal(t, result, 0)
}
