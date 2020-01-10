package ch16

import "testing"

import "github.com/magiconair/properties/assert"

func TestEnglishInt(t *testing.T) {
	ans := englishInt(123456789)
	assert.Equal(t, ans, "one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine ")

	ans = englishInt(100030000)
	assert.Equal(t, ans, "one hundred million thirty thousand ")

	ans = englishInt(-123456789)
	assert.Equal(t, ans, "negative one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine ")

}
