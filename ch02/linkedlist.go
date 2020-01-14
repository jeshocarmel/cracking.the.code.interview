package ch02

//Node denotes a node in the linked list
type Node struct {
	data int
	next *Node
}

func (n *Node) getAsList() []int {

	list := []int{}

	for n != nil {
		list = append(list, n.data)
		n = n.next
	}

	return list
}

func createLinkedList(list []int) *Node {

	head := &Node{data: list[0]}
	tmp := head

	for _, item := range list[1:] {
		tmp.next = &Node{data: item}
		tmp = tmp.next
	}
	return head
}

func (n *Node) len() int {
	head := n
	count := 0
	for head != nil {
		head = head.next
		count++
	}
	return count
}
