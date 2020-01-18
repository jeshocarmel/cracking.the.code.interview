package datastructures

//TreeNode represents a TreeNode in the binary search tree
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func createTreeNode(data int) *TreeNode {

	n := new(TreeNode)
	n.data = data
	return n
}

func (n *TreeNode) insert(data int) {
	if n == nil {
		return
	}

	if data < n.data {
		if n.left == nil {
			n.left = createTreeNode(data)
		} else {
			n.left.insert(data)
		}
	} else if data >= n.data {
		if n.right == nil {
			n.right = createTreeNode(data)
		} else {
			n.right.insert(data)
		}
	}
}

func (n *TreeNode) dfs(query int) (int, bool) {

	if n == nil {
		return 0, false
	}

	if n.data == query {
		return n.data, true
	} else if query < n.data {
		return n.left.dfs(query)
	} else {
		return n.right.dfs(query)
	}
}

func (n *TreeNode) bfs(query int) (int, bool) {

	if n == nil {
		return 0, false
	}

	// create a queue. BFS approach uses a queue
	queue := []*TreeNode{n}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		//search
		if current.data == query {
			return current.data, true
		}

		//add left node to queue
		if current.left != nil {
			queue = append(queue, current.left)
		}

		//add right node to queue
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return 0, false
}
