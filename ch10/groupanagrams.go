package ch10

import (
	"sort"
)

func groupAnagrams(arr []string) []string {

	mapList := make(map[string][]string)

	for _, str := range arr {
		sorted := sortChars(str)
		mapList[sorted] = append(mapList[sorted], str)
	}

	index := 0

	for key := range mapList {
		for _, val := range mapList[key] {
			arr[index] = val
			index++
		}
	}

	return arr

}

func sortChars(str string) string {

	r := []rune(str)
	sort.Sort(sortRunes(r))
	return string(r)

}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
