package ch01

func isUnique(str string) bool {

	mymap := make(map[rune]int)
	for _, ch := range str {
		mymap[ch]++
	}

	for _, v := range mymap {
		if v > 1 {
			return false
		}
	}

	return true
}

func isUniqueWithoutDataStructure(str string) bool {

	charset := make([]bool, 128)

	for _, ch := range str {
		if charset[ch] {
			return false
		}
		charset[ch] = true
	}

	return true
}
