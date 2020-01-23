package ch04

import "testing"

import "github.com/magiconair/properties/assert"

func TestIsBST(t *testing.T) {

	root := createNode(50)
	arr := []int{56, 27, 24, 39, 75, 95}

	//----case 1 starts here
	for _, item := range arr {
		insert(root, item)
	}

	/*
			56
		   /	\
		27		 75
		/	\		\
		24   39		95

	*/

	var result bool
	result = isBST(root)
	assert.Equal(t, result, true)

	//----case 1 ends

	//----case 2 starts here

	root.left.right = createNode(22)

	/*
			56
		   /	\
		27		 75
		/	\		\
		24   22		95

	*/
	result = isBST(root)
	assert.Equal(t, result, false)

}
