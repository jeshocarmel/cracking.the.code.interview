package datastructures

//Node represents the node of a double linked list
type Node struct {
	prev  *Node
	next  *Node
	Key   int
	Value int
}

//DoubleLinkedList represents the double linked list
type DoubleLinkedList struct {
	head *Node
	tail *Node
}

//Insert ...
func (dll *DoubleLinkedList) Insert(newNode *Node) *Node {

	if dll.head == nil {
		dll.head = newNode
		dll.tail = dll.head
	} else {
		dll.tail.next = newNode
		dll.tail = dll.tail.next
	}

	return newNode
}

//Insertafter ..
func (dll *DoubleLinkedList) Insertafter(node, newNode *Node) *Node {

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

//InsertBefore ...
func (dll *DoubleLinkedList) InsertBefore(node, newNode *Node) *Node {
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

//InsertAtBeginning ..
func (dll *DoubleLinkedList) InsertAtBeginning(newNode *Node) *Node {

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

//InsertAtEnd ..
func (dll *DoubleLinkedList) InsertAtEnd(newNode *Node) *Node {

	//handling empty list
	if dll.tail == nil {
		return dll.InsertAtBeginning(newNode)
	}

	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode

	return newNode
}

//Remove ..
func (dll *DoubleLinkedList) Remove(node *Node) *Node {

	for n := dll.head; n != nil; n = n.next {
		if node == n {
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

//GetLastNode ..
func (dll *DoubleLinkedList) GetLastNode() *Node {
	if dll.tail != nil {
		return dll.tail
	}
	return nil
}

//GetKeysAsList ..
func (dll *DoubleLinkedList) GetKeysAsList() []int {

	list := []int{}

	for node := dll.head; node != nil; node = node.next {
		list = append(list, node.Key)
	}

	return list
}

//CreateNode ..
func CreateNode(k, v int) *Node {
	return &Node{Key: k, Value: v}
}
