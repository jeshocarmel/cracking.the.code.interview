package ch04

import (
	"fmt"
	"testing"
)

func TestListOfDepths(t *testing.T) {

	root := createNode(20)
	insert(root, 10)
	insert(root, 2)
	insert(root, 30)
	insert(root, 13)
	insert(root, 4)

	//get height of the tree
	height := getHeightOfTree(root)

	//create a slice of linked list with height as the length
	listarray := make([]*LinkedListNode, height)

	//populate the slice with level nodes info
	getListOfDepths(root, 0, listarray)

	for level, item := range listarray {

		fmt.Print("level ", level, ": ")
		for item != nil {
			fmt.Print(item.data, " ")
			item = item.next
		}

		fmt.Println()
	}

}
