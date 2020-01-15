package ch16

import "fmt"

//Grid denotes the infinite grid
type Grid [][]int

func (g Grid) create(size int) [][]int {

	g = make([][]int, size)

	for j := 0; j < size; j++ {
		g[j] = make([]int, size)
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if j+1 < len(g[0]) {
				if g[i][j] == 0 {
					g[i][j+1] = 1
				} else {
					g[i][j+1] = 0
				}
			}
		}
		if g[i][0] == 0 {
			if i+1 < len(g) {
				g[i+1][0] = 1
			}
		}
	}

	return g

}

func (g Grid) display() {

	for i := 0; i < len(g); i++ {
		fmt.Println(g[i])
	}
}

func printKMoves(k int) {

	fmt.Println(k)

	var antgrid Grid
	antgrid = antgrid.create(10)
	antgrid.display()

}
