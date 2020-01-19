package ch04

func (n *Node) commonAncestor(n1, n2 *Node) *Node {

	// if root is one of the nodes
	if n1.data == n.data || n2.data == n.data {
		return n
	} else if (n1.data < n.data && n2.data > n.data) || (n1.data > n.data && n2.data < n.data) { // if both the nodes are in different subtrees
		return n
	} else {
		if n1.data < n.data {
			return n.left.commonAncestor(n1, n2)
		} else {
			return n.right.commonAncestor(n1, n2)
		}

	}

	return nil
}

func (n *Node) getNode(data int) *Node {

	if n == nil {
		return nil
	}

	if n.data == data {
		return n
	}

	left := n.left.getNode(data)
	right := n.right.getNode(data)

	if left != nil {
		return left
	}

	return right
}
