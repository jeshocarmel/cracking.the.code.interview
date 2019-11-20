package ch10

import "fmt"

//GenerateMatrix generates a matrix
func GenerateMatrix(row int, col int) {

	mat := make([][]int, col)

	for i := range mat {
		mat[i] = make([]int, row)
	}

	fmt.Println(mat)

}
