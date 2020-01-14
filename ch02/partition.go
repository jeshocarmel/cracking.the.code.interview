package ch02

func partition(head *Node, x int) *Node {

	newhead := head
	lower := &Node{}
	higher := &Node{}

	for newhead != nil {

		if newhead.data < x {
			lower.data = newhead.data
			tmp := &Node{}
			tmp.next = lower
			lower = tmp
		} else {
			higher.data = newhead.data
			tmp := &Node{}
			tmp.next = higher
			higher = tmp
		}

		newhead = newhead.next
	}

	lower = lower.next
	higher = higher.next

	partioned := lower

	for lower.next != nil {
		lower = lower.next
	}

	lower.next = higher

	return partioned

}

func partition2(node *Node, x int) *Node {

	head := node
	tail := node

	for node != nil {

		next := node.next

		if node.data < x {
			node.next = head
			head = node
		} else {
			tail.next = node
			tail = node
		}
		node = next
	}

	tail.next = nil

	return head

}
