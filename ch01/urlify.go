package ch01

func urlify(str string, last int) string {

	strArray := []rune(str)

	last = last - 1

	for i := len(strArray) - 1; last > 0; {

		if strArray[last] != ' ' {
			strArray[i], strArray[last] = strArray[last], strArray[i]
			i--
		} else {
			strArray[i] = '0'
			strArray[i-1] = '2'
			strArray[i-2] = '%'
			i -= 3
		}
		last--
	}

	return string(strArray)
}
