package ch16

import "testing"

import "github.com/magiconair/properties/assert"

func TestNumberSwapperEasy(t *testing.T) {

	ans1, ans2 := numberSwapperEasy(5, 10)
	assert.Equal(t, 10, ans1)
	assert.Equal(t, 5, ans2)
}

func TestNumberSwapperComplex(t *testing.T) {
	ans1, ans2 := numberSwapperComplex(5, 10)
	assert.Equal(t, 10, ans1)
	assert.Equal(t, 5, ans2)
}
