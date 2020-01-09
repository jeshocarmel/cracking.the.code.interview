package advanced

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestNaiveSearch(t *testing.T) {

	indexes := naiveStringSearch("she sells sea shells on the seashore", "sea")
	expected := []int{10, 28}

	assert.Equal(t, indexes, expected)
}

func TestRabinKarp(t *testing.T) {

	indexes := rabinKarpSearch("she sells sea shells on the seashore", "sea")
	expected := []int{10, 28}

	assert.Equal(t, indexes, expected)
}
