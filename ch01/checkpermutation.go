package ch01

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortString(str string) string {

	r := []rune(str)
	sort.Sort(sortRunes(r))
	return string(r)

}

func checkPermutation1(str1, str2 string) bool {

	str1 = sortString(str1)
	str2 = sortString(str2)

	if str1 == str2 {
		return true
	}

	return false
}

func checkPermutation2(str1, str2 string) bool {

	charset := make([]int, 128)

	for _, ch := range str1 {
		charset[ch]++
	}

	for _, ch := range str2 {
		charset[ch]--
		if charset[ch] < 0 {
			return false
		}
	}

	return true

}
