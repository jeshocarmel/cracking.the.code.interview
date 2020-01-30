package ch16

import (
	"math"
	"sort"
)

//BruteForce - O(N^2)
func smallestDifferenceBruteForce(arr1, arr2 []int) int {

	smallest := math.MaxInt8

	for _, item1 := range arr1 {
		for _, item2 := range arr2 {
			if abs(item1-item2) < smallest {
				smallest = abs(item1 - item2)
			}
		}
	}

	return smallest
}

//Optimal O(m log m +n log n) where m and n are len(arr1) and len(arr2) as sorting takes place for both the arrays
func smallestDifference(arr1, arr2 []int) int {

	//sort arr1
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	//sort arr2
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	fp := 0 // pointer for arr1
	sp := 0 //pointer for arr2

	var smallest = math.MaxInt8

	for fp < len(arr1) && sp < len(arr2) {

		calculated := abs(arr1[fp] - arr2[sp])

		if calculated < smallest {
			smallest = calculated
		}

		if arr1[fp] < arr2[sp] {
			fp++
		} else {
			sp++
		}
	}
	return smallest
}
