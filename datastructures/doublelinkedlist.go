package ch16

//Node represents the node of a double linked list
type Node struct {
	prev *Node
	next *Node
	data int
}

//DoubleLinkedList represents the double linked list
type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoubleLinkedList) insert(i int) *Node {

	newNode := &Node{data: i}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = dll.head
	} else {
		dll.tail.next = newNode
		dll.tail = dll.tail.next
	}

	return newNode
}

func (dll *DoubleLinkedList) insertafter(i int, node *Node) *Node {

	newNode := &Node{data: i}
	newNode.prev = node

	if node.next == nil {
		node.next = newNode
		dll.tail = newNode

	} else {
		newNode.next = node.next
		node.next.prev = newNode
	}
	node.next = newNode

	return newNode
}

func (dll *DoubleLinkedList) insertBefore(i int, node *Node) *Node {
	newNode := &Node{data: i}
	newNode.next = node

	if node.prev == nil {
		node.prev = newNode
		dll.head = newNode
	} else {
		newNode.prev = node.prev
		node.prev.next = newNode
	}

	return newNode
}

func (dll *DoubleLinkedList) insertAtBeginning(i int) *Node {
	newNode := &Node{data: i}

	if dll.head == nil {
		dll.head = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}

	return newNode
}

func (dll *DoubleLinkedList) insertAtEnd(i int) *Node {
	newNode := &Node{data: i}

	//handling empty list
	if dll.tail == nil {
		return dll.insertAtBeginning(i)
	}

	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode

	return newNode
}

func (dll *DoubleLinkedList) getAsList() []int {

	list := []int{}

	for node := dll.head; node != nil; node = node.next {
		list = append(list, node.data)
	}

	return list
}
