package ch03

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSortStack(t *testing.T) {

	s := newSortStack()

	s.push(10)
	s.push(9)
	s.pop()
	s.push(11)
	s.push(19)
	s.push(13)

	s.sort()

	expected := []int{19, 13, 11, 10}
	assert.Equal(t, s.items, expected)
}
