package ch01

func rotateMatrix(mat [][]int) [][]int {

	n := len(mat)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			mat[j][n-1-i], mat[i][j], mat[n-1-j][i], mat[n-1-i][n-1-j] = mat[i][j], mat[n-1-j][i], mat[n-1-i][n-1-j], mat[j][n-1-i]
		}
	}

	return mat
}
