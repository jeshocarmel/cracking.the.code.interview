package datastructures

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {

	g := new(Graph)

	A := g.CreateGraphNode("A")
	B := g.CreateGraphNode("B")
	C := g.CreateGraphNode("C")
	D := g.CreateGraphNode("D")
	F := g.CreateGraphNode("F")
	E := g.CreateGraphNode("E")

	addEdge(A, B)
	addEdge(A, C)
	addEdge(A, D)
	addEdge(B, C)
	addEdge(B, E)
	addEdge(E, B)
	addEdge(B, F)

	fmt.Print("BFS: ")
	bfs(A)
	fmt.Println("")

	// clearing the visited status after bfs
	g.clearVisits()

	fmt.Print("DFS: ")
	dfs(A)
	fmt.Println("")
}
