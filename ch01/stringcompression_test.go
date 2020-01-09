package ch01

import "testing"

import "github.com/magiconair/properties/assert"

func TestCompressString(t *testing.T) {
	ans := compressString("aabcccccaaa")
	assert.Equal(t, ans, "a2b1c5a3")
}
