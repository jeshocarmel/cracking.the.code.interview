package ch16

import "testing"

import "github.com/magiconair/properties/assert"

func TestFactorialZeros(t *testing.T) {
	ans := countZeros(5)
	assert.Equal(t, ans, 1)

	ans = countZeros(11)
	assert.Equal(t, ans, 2)
}
