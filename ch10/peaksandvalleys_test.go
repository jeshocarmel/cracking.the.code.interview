package ch10

import (
	"sort"
	"testing"
)

func TestPeaksAndValleys(t *testing.T) {

	arr := generateArray(10)
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
