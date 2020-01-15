package ch16

import (
	"cracking.the.code.interview/datastructures"
)

var cache map[int]int
var capacity int
var dll datastructures.DoubleLinkedList

func init() {
	cache = make(map[int]int)
}

func createCache(c int) map[int]int {
	capacity = c
	return cache
}

func addToCache(key, value int) {

	if _, ok := cache[key]; ok {
		removeFromCache(key)
	}

	if len(cache) >= capacity {
		lastaccessedkey := dll.GetLastKey()
		dll.Remove(lastaccessedkey)
		delete(cache, lastaccessedkey)
	}

	dll.InsertAtBeginning(key, value)
	cache[key] = value

}

func removeFromCache(key int) {

	delete(cache, key)
	dll.Remove(key)

}
