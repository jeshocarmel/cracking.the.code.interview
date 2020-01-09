package ch01

import "testing"
import "github.com/magiconair/properties/assert"

func TestIsUnique(t *testing.T) {

	ans := isUnique("alphabet")
	assert.Equal(t, false, ans)

	ans = isUnique("uncopyrightables")
	assert.Equal(t, true, ans)

}

func TestIsUniqueWithoutDataStructure(t *testing.T) {
	ans := isUniqueWithoutDataStructure("alphabet")
	assert.Equal(t, false, ans)

	ans = isUniqueWithoutDataStructure("uncopyrightables")
	assert.Equal(t, true, ans)
}
