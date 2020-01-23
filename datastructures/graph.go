package datastructures

import "fmt"

//GraphNode represents a node in the graph
type GraphNode struct {
	data     string
	children []*GraphNode
	visited  bool
}

//Graph represents the datastructure as a whole
type Graph struct {
	nodes []*GraphNode
}

func (g *Graph) createNode(data string) *GraphNode {

	n := &GraphNode{data: data}
	n.children = make([]*GraphNode, 0)

	g.nodes = append(g.nodes, n)

	return n
}

func addEdge(src, target *GraphNode) {
	src.children = append(src.children, target)
}

func dfs(root *GraphNode) {

	if root == nil {
		return
	}

	fmt.Print(root.data, " ")
	root.visited = true

	for _, child := range root.children {

		if !child.visited {
			dfs(child)
		}
	}

}

func bfs(root *GraphNode) {

	if root == nil {
		return
	}

	root.visited = true
	queue := []*GraphNode{}
	queue = append(queue, root)

	for len(queue) != 0 {

		GraphNode := queue[0]
		fmt.Print(GraphNode.data, " ")
		queue = queue[1:]

		for _, child := range GraphNode.children {
			if child.visited == false {
				child.visited = true
				queue = append(queue, child)
			}
		}
	}
}

func (g *Graph) clearVisits() {

	for _, node := range g.nodes {
		node.visited = false
	}
}
