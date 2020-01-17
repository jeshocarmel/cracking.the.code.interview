package ch03

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestStackOfPlates(t *testing.T) {

	myplates := newSetOfPlates(3)

	myplates.push(1)
	myplates.push(2)
	myplates.push(3)
	myplates.pop()
	myplates.push(4)
	myplates.push(5)
	myplates.pop()
	myplates.push(6)
	myplates.push(7)
	myplates.push(8)
	myplates.push(9)
	myplates.push(10)
	myplates.push(11)
	myplates.push(12)

	assert.Equal(t, myplates.items, [][]int{{1, 2, 4}, {6, 7, 8}, {9, 10, 11}, {12}})

	myplates.pop()
	myplates.pop()
	myplates.pop()

	assert.Equal(t, myplates.items, [][]int{{1, 2, 4}, {6, 7, 8}, {9}})

	popped := myplates.popArray(1)
	assert.Equal(t, popped, 8)
	assert.Equal(t, myplates.items, [][]int{{1, 2, 4}, {6, 7}, {9}})
}
