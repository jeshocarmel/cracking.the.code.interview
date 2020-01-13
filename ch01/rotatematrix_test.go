package ch01

import "testing"

import "github.com/magiconair/properties/assert"

func TestRotateMatrix(t *testing.T) {

	mat := make([][]int, 4)
	mat[0] = []int{1, 2, 3, 4}
	mat[1] = []int{5, 6, 7, 8}
	mat[2] = []int{9, 0, 1, 2}
	mat[3] = []int{3, 4, 5, 6}

	ans := rotateMatrix(mat)

	expected := make([][]int, 4)

	expected[0] = []int{3, 9, 5, 1}
	expected[1] = []int{4, 0, 6, 2}
	expected[2] = []int{5, 1, 7, 3}
	expected[3] = []int{6, 2, 8, 4}

	assert.Equal(t, ans, expected)

}
