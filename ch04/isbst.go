package ch04

func isBST(root *Node, min, max interface{}) bool {
	if root == nil {
		return true
	}

	if min != nil && root.data < min.(int) {
		return false
	}

	if max != nil && root.data > max.(int) {
		return false
	}

	return isBST(root.left, min, root.data) && isBST(root.right, root.data, max)

}
