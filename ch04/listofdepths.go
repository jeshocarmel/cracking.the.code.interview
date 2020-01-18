package ch04

func (n *Node) insert(data int) {

	if n.data < data {
		if n.left == nil {
			n.left = createNode(data)
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = createNode(data)
		} else {
			n.right.insert(data)
		}
	}
}

func (n *Node) dfsTraversal(mymap map[int][]int, depth int) {

	if n == nil {
		return
	}

	mymap[depth] = append(mymap[depth], n.data)

	n.left.dfsTraversal(mymap, depth+1)
	n.right.dfsTraversal(mymap, depth+1)
}

func (n *Node) bfsTraversal(mymap map[int][]int) {

	queue := []*Node{n}
	depth := 0

	for len(queue) != 0 {

		parents := queue
		queue = []*Node{}

		for _, parent := range parents {

			mymap[depth] = append(mymap[depth], parent.data)

			if parent.left != nil {
				queue = append(queue, parent.left)
			}
			if parent.right != nil {
				queue = append(queue, parent.right)
			}
		}
		depth++
	}
}
