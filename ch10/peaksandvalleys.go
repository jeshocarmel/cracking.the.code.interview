package ch10

func alternatePeaksAndValleys(arr []int) []int {

	for i := 1; i < len(arr); i += 2 {
		arr[i-1], arr[i] = arr[i], arr[i-1]
	}

	return arr
}

func alternatePeaksAndValleysOptimal(arr []int) []int {

	return arr
}
