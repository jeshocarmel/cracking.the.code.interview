package ch04

import "testing"
import "github.com/stretchr/testify/assert"

var g Graph

func TestRouteBetweenNodesDfs(t *testing.T) {

	g := new(Graph)

	A := g.CreateGraphNode("A")
	B := g.CreateGraphNode("B")
	C := g.CreateGraphNode("C")
	D := g.CreateGraphNode("D")
	N1 := g.CreateGraphNode("N1")
	N2 := g.CreateGraphNode("N2")

	addEdge(A, B)
	addEdge(A, C)
	addEdge(A, D)
	addEdge(B, C)
	addEdge(N1, N2)

	// N1 and N2 are independent to the rest of the node

	/*
		A --- B
		| \
		|  \
		C   D		N1 --- N2
	*/

	var routeExists bool
	routeExists = dfs(A, N2)
	assert.Equal(t, routeExists, false)

	g.clearVisits()

	addEdge(D, N1) //N1 is connected to D

	/*
		A --- B
		| \
		|  \
		C   D --- N1 --- N2
	*/

	routeExists = dfs(A, N2)
	assert.Equal(t, routeExists, true)
}

func TestRouteBetweenNodesBfs(t *testing.T) {

	g := new(Graph)

	A := g.CreateGraphNode("A")
	B := g.CreateGraphNode("B")
	C := g.CreateGraphNode("C")
	D := g.CreateGraphNode("D")
	N1 := g.CreateGraphNode("N1")
	N2 := g.CreateGraphNode("N2")

	// N1 and N2 are independent to the rest of the nodes

	/*
		A --- B
		| \
		|  \
		C   D		N1 --- N2
	*/

	addEdge(A, B)
	addEdge(A, C)
	addEdge(A, D)
	addEdge(B, C)
	addEdge(N1, N2)

	var routeExists bool
	routeExists = bfs(A, N2)
	assert.Equal(t, routeExists, false)

	g.clearVisits()

	/*
		A --- B
		| \
		|  \
		C   D --- N1 --- N2
	*/

	addEdge(D, N1) //N1 is connected to D
	routeExists = bfs(A, N2)
	assert.Equal(t, routeExists, true)

}
