package ch01

import "testing"

import "github.com/magiconair/properties/assert"

func TestUrlify(t *testing.T) {

	str := "Mr John Smith    "
	ans := urlify(str, 13)
	assert.Equal(t, ans, "Mr%20John%20Smith")
}
