package ch02

func detectLoop(node *Node) *Node {

	slow := node
	fast := node

	for slow != nil && fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			break
		}
	}

	if fast == nil || slow == nil {
		return nil
	}

	slow = node
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return fast

}
