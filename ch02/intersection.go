package ch02

func intersection(node1, node2 *Node) (*Node, bool) {

	tail1 := node1
	tail2 := node2

	// check whether intersection exists
	for tail1.next != nil {
		tail1 = tail1.next
	}

	for tail2.next != nil {
		tail2 = tail2.next
	}

	if tail1 != tail2 {
		return nil, false
	}

	// adjust difference in length for the two lists
	len1 := node1.len()
	len2 := node2.len()

	diff := abs(len1 - len2)

	if len1 > len2 {
		for i := 0; i < diff; i++ {
			node1 = node1.next
		}
	} else {
		for i := 0; i < diff; i++ {
			node2 = node2.next
		}
	}

	// finding the intersection
	for node1 != node2 {

		node1 = node1.next
		node2 = node2.next
	}

	return node1, true

}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
