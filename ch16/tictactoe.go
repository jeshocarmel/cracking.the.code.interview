package ch16

func hasWonBruteForce(arr [][]int) bool {

	if diagonalcheck(arr) || crossdiagonalcheck(arr) {
		return true
	}

	for i := 0; i < len(arr); i++ {
		if horizontalcheck(arr, i) {
			return true
		}

		if verticalcheck(arr, i) {
			return true
		}
	}

	return false

}

func horizontalcheck(arr [][]int, row int) bool {

	j := 0
	for j < len(arr[0]) {
		if arr[row][j] == 0 {
			return false
		}
		j++
	}

	return true
}

func verticalcheck(arr [][]int, col int) bool {

	i := 0

	for i < len(arr) {
		if arr[i][col] == 0 {
			return false
		}
		i++
	}

	return true

}

func diagonalcheck(arr [][]int) bool {

	for i, j := 0, 0; i < len(arr) && j < len(arr[0]); i, j = i+1, j+1 {
		if arr[i][j] == 0 {
			return false
		}
	}

	return true
}

func crossdiagonalcheck(arr [][]int) bool {

	for i, j := len(arr)-1, 0; i >= 0 && j < len(arr[0]); i, j = i-1, j+1 {
		if arr[i][j] == 0 {
			return false
		}
	}

	return true
}
