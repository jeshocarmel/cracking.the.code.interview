package ch16

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSmallestDifference(t *testing.T) {

	arr1 := []int{1, 2, 11, 5}
	arr2 := []int{4, 12, 19, 23, 127, 235}

	difference := smallestDifferenceBruteForce(arr1, arr2)
	assert.Equal(t, difference, 1)

	difference = smallestDifference(arr1, arr2)
	assert.Equal(t, difference, 1)

}
