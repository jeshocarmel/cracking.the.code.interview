package ch10

func alternatePeaksAndValleys(arr []int) []int {

	for i := 1; i < len(arr); i += 2 {
		arr[i-1], arr[i] = arr[i], arr[i-1]
	}

	return arr
}

func alternatePeaksAndValleysOptimal(arr []int) []int {

	for i := 1; i < len(arr); i += 2 {
		biggestindex := findBiggestIndex(arr, i-1, i, i+1)
		if i != biggestindex {
			arr[biggestindex], arr[i] = arr[i], arr[biggestindex]
		}
	}

	return arr
}

func findBiggestIndex(arr []int, a, b, c int) int {

	aVal := -1
	bVal := -1
	cVal := -1

	arrSize := len(arr)

	if a < arrSize {
		aVal = arr[a]
	}
	if b < arrSize {
		bVal = arr[b]
	}
	if c < arrSize {
		cVal = arr[c]
	}

	max := Max(aVal, Max(bVal, cVal))

	if aVal == max {
		return a
	} else if bVal == max {
		return b
	} else {
		return c
	}
}
