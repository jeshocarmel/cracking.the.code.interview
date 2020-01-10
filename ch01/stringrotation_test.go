package ch01

import "testing"

import "github.com/magiconair/properties/assert"

func TestIsSubstring(t *testing.T) {
	ans := isSubString("erbottlewat", "waterbottle")
	assert.Equal(t, ans, true)

	ans = isSubString("robertjeshoveloncarmel", "elrob")
	assert.Equal(t, ans, true)

}
