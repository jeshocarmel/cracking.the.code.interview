package ch10

import (
	"math/rand"
	"time"
)

//GenerateMatrix generates a matrix
func GenerateMatrix(row int, col int) [][]int {

	mat := make([][]int, col)

	for i := range mat {
		mat[i] = make([]int, row)
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(50)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {

			if i == 0 && j == 0 {
				mat[i][j] = random
			} else if j == 0 {
				mat[i][j] = mat[i-1][j] + 5
			} else if j > 0 {
				mat[i][j] = mat[i][j-1] + 5
			}

		}
	}

	return mat
}

//SearchMatrix searches for an item in the sorted matrix
func SearchMatrix(mat [][]int, x int) bool {

	row := 0
	col := len(mat[0]) - 1

	for col >= 0 && row < len(mat) {
		if mat[row][col] == x {
			return true
		} else if mat[row][col] > x {
			col--
		} else {
			row++
		}
	}

	return false
}
