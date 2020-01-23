package datastructures

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {

	g := new(Graph)

	A := g.createNode("A")
	B := g.createNode("B")
	C := g.createNode("C")
	D := g.createNode("D")
	F := g.createNode("F")
	E := g.createNode("E")

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
