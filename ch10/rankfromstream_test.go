package ch10

import (
	"testing"
)

func TestGetRank(t *testing.T) {

	// generate a array of random numbers
	arr := []int{20, 15, 25, 10, 23, 5, 13, 24}

	first := arr[0]
	arr = arr[1:]

	// create root element
	rn := &RankNode{data: first}

	// insert other elements into the tree
	for _, item := range arr {
		rn.insert(item)
	}

	//display tree by inorder traversal
	//inorder(rn)

	if rn.getRank(24) != 6 {
		t.Error("wrong output")
		t.Fail()
	}

}
