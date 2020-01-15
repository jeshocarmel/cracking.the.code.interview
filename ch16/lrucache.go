package ch16

import (
	"cracking.the.code.interview/datastructures"
)

var cache map[int]*datastructures.Node
var capacity int
var dll datastructures.DoubleLinkedList

func init() {
	cache = make(map[int]*datastructures.Node)
}

func createCache(c int) map[int]*datastructures.Node {
	capacity = c
	return cache
}

func addToCache(key, value int) {

	if _, ok := cache[key]; ok {
		removeFromCache(key)
	}

	if len(cache) >= capacity {
		lastNode := dll.GetLastNode()
		dll.Remove(lastNode)
		delete(cache, lastNode.Key)
	}

	newnode := datastructures.CreateNode(key, value)
	dll.InsertAtBeginning(newnode)
	cache[key] = newnode

}

func removeFromCache(key int) {

	dll.Remove(cache[key])
	delete(cache, key)

}

func getKeysOfCache() []int {

	keys := []int{}

	for k := range cache {
		keys = append(keys, k)
	}

	return keys
}
