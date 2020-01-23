package ch04

//LinkedListNode represents a node in a linked list
type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func createLinkedListNode(data int) *LinkedListNode {
	return &LinkedListNode{data: data}
}

func getListOfDepths(root *Node, level int, listArray []*LinkedListNode) {

	if root == nil {
		return
	}

	if listArray[level] == nil {
		listArray[level] = createLinkedListNode(root.data)
	} else {
		lnode := listArray[level]
		for lnode.next != nil {
			lnode = lnode.next
		}
		lnode.next = createLinkedListNode(root.data)
	}

	getListOfDepths(root.left, level+1, listArray)
	getListOfDepths(root.right, level+1, listArray)

}

func getHeightOfTree(root *Node) int {

	if root == nil {
		return 0
	}

	lHeight := getHeightOfTree(root.left)
	rHeight := getHeightOfTree(root.right)

	if lHeight > rHeight {
		return lHeight + 1
	}

	return rHeight + 1
}
