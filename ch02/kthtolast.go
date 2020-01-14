package ch02

func kthToLast(head *Node, k int) int {

	fast := head
	slow := head

	n := 1

	for fast.next != nil {
		fast = fast.next
		n++
	}

	// the position of kth to last will be at i
	pos := n - k
	i := 1

	for i != pos {
		slow = slow.next
		i++
	}

	return slow.data
}
