package ch16

func numberSwapperEasy(a, b int) (int, int) {

	a, b = b, a
	return a, b
}

func numberSwapperComplex(a, b int) (int, int) {

	b = a + b
	a = b - a
	b = b - a

	return a, b
}
