package ch16

import "testing"

import "github.com/magiconair/properties/assert"

func TestTicTacToe(t *testing.T) {

	/* case 1

	x o x
	o x o
	x o x

	winner - x
	*/

	arr := [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	result := hasWon(arr, 2, 0)
	assert.Equal(t, result, 1)

	/* case 2

	x o o
	o x x
	x x o

	result = no winner
	*/

	arr = [][]int{{1, 0, 0}, {0, 1, 1}, {1, 1, 0}}
	result = hasWon(arr, 0, 2)
	assert.Equal(t, result, -1)
}
