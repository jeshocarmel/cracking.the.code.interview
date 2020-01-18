package ch04

//Node represents the node of a binary tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

func createNode(data int) *Node {
	return &Node{data: data}
}

func createMinimalBST(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2

	node := createNode(arr[mid])
	node.left = createMinimalBST(arr[0:mid])
	node.right = createMinimalBST(arr[mid+1:])
	return node
}

func (n *Node) maxdepth() int {

	if n == nil {
		return 0
	}

	ldepth := n.left.maxdepth()
	rdepth := n.right.maxdepth()

	if ldepth > rdepth {
		return ldepth + 1
	}

	return rdepth + 1

}
