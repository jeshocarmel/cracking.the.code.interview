package ch01

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestZeroMatrix(t *testing.T) {

	newmat := [][]int{
		[]int{1, 2, 5, 9},
		[]int{1, 7, 5, 8},
		[]int{4, 6, 0, 1}}

	ans := zeroMatrix(newmat)

	expected := [][]int{
		[]int{1, 2, 0, 9},
		[]int{1, 7, 0, 8},
		[]int{0, 0, 0, 0}}

	ans = zeroMatrixSpaceOptimized(newmat)

	assert.Equal(t, ans, expected)

}
