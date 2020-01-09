package ch01

import (
	"strings"
)

func isPalindromePermutation(str string) bool {

	countmap := make(map[rune]int)

	for _, ch := range strings.ReplaceAll(strings.ToLower(str), " ", "") {
		countmap[ch]++
	}

	odd := 0
	for _, v := range countmap {
		if v%2 != 0 {
			odd++
			if odd > 1 {
				return false
			}
		}
	}
	return true
}

func isPalindromePermutationSimplified(str string) bool {

	//credit : https://github.com/krlv/ctci-go/blob/master/ch01/04_palindrome_permutation.go

	str = strings.ReplaceAll(strings.ToLower(str), " ", "")

	countmap := make(map[rune]int)

	odd := 0

	for _, ch := range str {
		countmap[ch]++

		if countmap[ch]%2 != 0 {
			odd++
		} else {
			odd--
		}
	}

	return !(odd > 1)

}
