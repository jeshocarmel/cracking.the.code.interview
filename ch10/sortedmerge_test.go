package ch10

import (
	"testing"
)

func TestSortedMerge(t *testing.T) {

	arr1 := []int{17, 22, 34, 41, 56, 0, 0, 0}
	arr2 := []int{61, 79, 82}

	merge(arr1, arr2, 5, 3)

	expected := []int{17, 22, 34, 41, 56, 61, 79, 82}
	if !isEqualIntSlices(arr1, expected) {
		t.Fail()
	}
}

func isEqualIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
