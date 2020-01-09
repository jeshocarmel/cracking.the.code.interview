package ch01

import "testing"

import "github.com/magiconair/properties/assert"

func TestCheckPermutation1(t *testing.T) {

	ans := checkPermutation1("dormitory", "dirtyroom")
	assert.Equal(t, ans, true)
	ans = checkPermutation1("conversation", "voicesranton")
	assert.Equal(t, ans, true)
	ans = checkPermutation1("ringostarr", "johnlennon")
	assert.Equal(t, ans, false)

}

func TestCheckPermutation2(t *testing.T) {

	ans := checkPermutation2("dormitory", "dirtyroom")
	assert.Equal(t, ans, true)
	ans = checkPermutation2("conversation", "voicesranton")
	assert.Equal(t, ans, true)
	ans = checkPermutation2("ringostarr", "johnlennon")
	assert.Equal(t, ans, false)

}
