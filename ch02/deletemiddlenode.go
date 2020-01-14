package ch02

func deleteMiddleNode(n *Node) {

	if n == nil || n.next == nil {
		return
	}

	n.data = n.next.data
	n.next = n.next.next

}
