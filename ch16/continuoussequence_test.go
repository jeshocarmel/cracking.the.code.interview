package ch16

import "testing"

import "github.com/magiconair/properties/assert"

func TestFindContinuousSequence(t *testing.T) {

	ans := findContinuousSequence([]int{2, -8, 3, -2, 4, -10})
	assert.Equal(t, ans, 5)

	ans = findContinuousSequence([]int{2, -8, 3, -2, 4, -1, 3})
	assert.Equal(t, ans, 7)
}
