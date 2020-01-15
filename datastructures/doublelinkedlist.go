package datastructures

//Node represents the node of a double linked list
type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

//DoubleLinkedList represents the double linked list
type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoubleLinkedList) Insert(k int, v int) *Node {

	newNode := &Node{key: k, value: v}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = dll.head
	} else {
		dll.tail.next = newNode
		dll.tail = dll.tail.next
	}

	return newNode
}

func (dll *DoubleLinkedList) Insertafter(k, v int, node *Node) *Node {

	newNode := &Node{key: k, value: v}
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

func (dll *DoubleLinkedList) InsertBefore(k, v int, node *Node) *Node {
	newNode := &Node{key: k, value: v}
	newNode.next = node

	if node.prev == nil {
		node.prev = newNode
		dll.head = newNode
	} else {
		newNode.prev = node.prev
		node.prev.next = newNode
	}
	node.prev = newNode

	return newNode
}

func (dll *DoubleLinkedList) InsertAtBeginning(k, v int) *Node {
	newNode := &Node{key: k, value: v}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}

	return newNode
}

func (dll *DoubleLinkedList) InsertAtEnd(k, v int) *Node {
	newNode := &Node{key: k, value: v}

	//handling empty list
	if dll.tail == nil {
		return dll.InsertAtBeginning(k, v)
	}

	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode

	return newNode
}

func (dll *DoubleLinkedList) Remove(key int) *Node {

	for n := dll.head; n != nil; n = n.next {
		if key == n.key {
			if n.prev == nil {
				dll.head = n.next
			} else {
				n.prev.next = n.next
			}

			if n.next == nil {
				dll.tail = n.prev
			} else {
				n.next.prev = n.prev
			}

			return n
		}
	}

	return nil

}

func (dll *DoubleLinkedList) GetLastKey() int {
	if dll.tail != nil {
		return dll.tail.key
	}
	return 0
}

func (dll *DoubleLinkedList) GetKeysAsList() []int {

	list := []int{}

	for node := dll.head; node != nil; node = node.next {
		list = append(list, node.key)
	}

	return list
}
