package ch16

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestLRUCache(t *testing.T) {

	cache = createCache(5)

	for i := 1; i <= 20; i++ {
		addToCache(i, i*100)
	}

	addToCache(1, 1234)

	expected := []int{1, 17, 18, 19, 20}
	assert.Equal(t, getKeysOfCache(), expected)

}
