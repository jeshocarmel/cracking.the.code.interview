package ch10

import (
	"sort"
	"testing"
)

func TestPeaksAndValleys(t *testing.T) {

	// to generate a random array use this
	//arr := generateArray(10)

	arr := []int{9, 1, 0, 4, 8, 7}
	sort.Ints(arr)

	peaksandvalleys := alternatePeaksAndValleys(arr)
	//fmt.Println(peaksandvalleys)

	for i := 2; i < len(peaksandvalleys); i += 2 {
		if (peaksandvalleys[i-2] < peaksandvalleys[i-1]) || (peaksandvalleys[i-1] > peaksandvalleys[i]) {
			t.Error(peaksandvalleys[i-2], peaksandvalleys[i-1], peaksandvalleys[i])
			t.Fail()
		}
	}

}

func TestPeaksAndValleysOptimal(t *testing.T) {

	// to generate a random array use this
	//arr := generateArray(10)

	arr := []int{9, 1, 0, 4, 8, 7}
	peaksandvalleys := alternatePeaksAndValleysOptimal(arr)
	//fmt.Println(peaksandvalleys)

	for i := 2; i < len(peaksandvalleys); i += 2 {
		if (peaksandvalleys[i-2] > peaksandvalleys[i-1]) || (peaksandvalleys[i-1] < peaksandvalleys[i]) {
			t.Error(peaksandvalleys[i-2], peaksandvalleys[i-1], peaksandvalleys[i])
			t.Fail()
		}
	}

}
