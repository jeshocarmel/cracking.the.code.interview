package ch16

func hasWon(arr [][]int, row int, col int) int {

	if hasWonByRow(arr, row) || hasWonByCol(arr, col) || hasWonByDiagonal(arr, row, col) {
		return arr[row][col]
	}

	return -1
}

func hasWonByRow(arr [][]int, row int) bool {

	for j := 1; j < len(arr[0]); j++ {
		if arr[row][j] != arr[row][0] {
			return false
		}
	}

	return true
}

func hasWonByCol(arr [][]int, col int) bool {

	for i := 1; i < len(arr); i++ {
		if arr[i][col] != arr[0][col] {
			return false
		}
	}

	return true
}

func hasWonByDiagonal(arr [][]int, row, col int) bool {

	if row == col {
		for i, j := 1, 1; i < len(arr) && j < len(arr[0]); i, j = i+1, j+1 {
			if arr[i][j] != arr[0][0] {
				return false
			}
		}
		return true
	} else {

		for i, j := len(arr)-2, 1; i >= 0 && j < len(arr[0]); i, j = i-1, j+1 {
			if arr[i][j] != arr[len(arr)-1][0] {
				return false
			}
		}
		return true
	}
}
