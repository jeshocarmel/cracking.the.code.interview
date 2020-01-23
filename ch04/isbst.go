package ch04

func isBST(root *Node) bool {
	if root == nil {
		return true
	}

	if root.left != nil && root.left.data > root.data {
		return false
	}

	if root.right != nil && root.right.data < root.data {
		return false
	}

	return isBST(root.left) && isBST(root.right)

}
