package ch01

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestOneAway(t *testing.T) {
	ans := isOneAway("pale", "ple")
	assert.Equal(t, ans, true)
	ans = isOneAway("pale", "pile")
	assert.Equal(t, ans, true)
	ans = isOneAway("pile", "piles") //haha
	assert.Equal(t, ans, true)
	ans = isOneAway("pile", "plea")
	assert.Equal(t, ans, false)
}
