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

	expected := make(map[int]int)
	expected[17] = 1700
	expected[18] = 1800
	expected[19] = 1900
	expected[20] = 2000
	expected[1] = 1234

	assert.Equal(t, expected, cache)

}
