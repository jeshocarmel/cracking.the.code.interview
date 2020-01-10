package ch01

func zeroMatrix(mat [][]int) [][]int {

	rowsToZero := []int{}
	colsToZero := []int{}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				rowsToZero = append(rowsToZero, i)
				colsToZero = append(colsToZero, j)
			}
		}
	}

	newmat := make([][]int, len(mat))

	for i := 0; i < len(mat); i++ {
		newmat[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			if find(rowsToZero, i) || find(colsToZero, j) {
				newmat[i][j] = 0
			} else {
				newmat[i][j] = mat[i][j]
			}
		}
	}

	return newmat
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func zeroMatrixSpaceOptimized(mat [][]int) [][]int {

	rowHasZero := false
	colHasZero := false

	// identify if first row has zero
	for j := 0; j < len(mat[0]); j++ {
		if mat[0][j] == 0 {
			rowHasZero = true
			break
		}
	}

	// identify if first col has zero
	for i := 0; i < len(mat); i++ {
		if mat[i][0] == 0 {
			colHasZero = true
			break
		}
	}

	// mark the first row or col as zero if any of it's internals has zero
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

	//nullify rows based on the first column
	for i := 1; i < len(mat); i++ {
		if mat[i][0] == 0 {
			nullifyRow(mat, i)
		}
	}

	//nullify cols based on the first row
	for j := 1; j < len(mat[0]); j++ {
		if mat[0][j] == 0 {
			nullifyCol(mat, j)
		}
	}

	if rowHasZero {
		nullifyRow(mat, 0)
	}

	if colHasZero {
		nullifyCol(mat, 0)
	}

	return mat

}

func nullifyRow(mat [][]int, row int) {

	for j := 0; j < len(mat[0]); j++ {
		mat[row][j] = 0
	}

}

func nullifyCol(mat [][]int, col int) {

	for i := 0; i < len(mat); i++ {
		mat[i][col] = 0
	}
}
