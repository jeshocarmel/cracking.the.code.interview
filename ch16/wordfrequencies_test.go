package ch16

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestWordFrequencies(t *testing.T) {

	mybook := []string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}

	parseBook(mybook)
	ans := getWordCount("brown")
	assert.Equal(t, ans, 1)

	ans = getWordCount("the")
	assert.Equal(t, ans, 2)
}
