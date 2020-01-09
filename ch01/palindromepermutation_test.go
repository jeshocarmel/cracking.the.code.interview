package ch01

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsPalindromePermutation(t *testing.T) {

	ans := isPalindromePermutation("Tact Coa")
	assert.Equal(t, ans, true)

	ans = isPalindromePermutation("Able was I ere I saw Elba")
	assert.Equal(t, ans, true)

	ans = isPalindromePermutation("beatles")
	assert.Equal(t, ans, false)

}

func TestIsPalindromePermutationSimplified(t *testing.T) {

	ans := isPalindromePermutationSimplified("Tact Coa")
	assert.Equal(t, ans, true)

	ans = isPalindromePermutationSimplified("Able was I ere I saw Elba")
	assert.Equal(t, ans, true)

	ans = isPalindromePermutationSimplified("beatles")
	assert.Equal(t, ans, false)

}
