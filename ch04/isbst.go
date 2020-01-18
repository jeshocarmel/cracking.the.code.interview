package ch04

var lastVisited interface{}

func (n *Node) isBST() bool {

	if n == nil {
		return true
	}
	if !n.left.isBST() {
		return false
	}
	if lastVisited != nil && n.data <= lastVisited.(int) {
		return false
	}

	lastVisited = n.data

	if !n.right.isBST() {
		return false
	}
	return true
}

func (n *Node) isBSTRecursive(min, max interface{}) bool {

	if n == nil {
		return true
	}

	n.left.isBSTRecursive(n.data, nil)

	if (min != nil && n.data > min.(int)) || (max != nil && n.data < max.(int)) {
		return false
	}

	n.right.isBSTRecursive(nil, n.data)
	return true
}
