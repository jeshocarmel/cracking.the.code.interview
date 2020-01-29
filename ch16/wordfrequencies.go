package ch16

import "strings"

var dictionary = make(map[string]int)

func parseBook(book []string) {

	for _, word := range book {
		word = strings.ToLower(word)
		dictionary[word]++
	}
}

func getWordCount(word string) int {

	word = strings.ToLower(word)

	if val, ok := dictionary[word]; ok {
		return val
	}

	return 0
}
