package ch01

func isOneAway(str1, str2 string) bool {

	d := len(str2) - len(str1)

	switch d {
	case 0:
		return isOneReplaced(str1, str2)
	case 1:
		return isOneInserted(str1, str2)
	case -1:
		return isOneInserted(str2, str1)
	default:
		return false
	}

}

func isOneInserted(str1, str2 string) bool {

	for i1, i2 := 0, 0; i2 < len(str1); {
		if str1[i1] != str2[i2] {
			i2++
		} else {
			i1++
			i2++
		}

		if i2-i1 > 1 {
			return false
		}
	}
	return true
}

func isOneReplaced(str1, str2 string) bool {

	found := false

	for i := range str1 {
		if str1[i] != str2[i] {
			if found {
				return false
			}
			found = true
		}
	}

	return found
}

func abs(i int) int {

	if i < 0 {
		return -i
	}
	return i
}
