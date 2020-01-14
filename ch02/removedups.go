package ch02

func removeDups(head *Node) *Node {

	counter := make(map[int]bool)

	newhead := head
	prev := head

	for newhead != nil {
		if counter[newhead.data] {
			prev.next = newhead.next
		} else {
			counter[newhead.data] = true
			prev = newhead
		}
		newhead = newhead.next
	}

	return head
}

func removeDupsOptimized(head *Node) *Node {

	for slow := head; slow != nil; slow = slow.next {
		prev := slow

		for fast := slow.next; fast != nil; fast = fast.next {
			if slow.data == fast.data {
				prev.next = fast.next
			} else {
				prev = fast
			}
		}
	}

	return head
}
