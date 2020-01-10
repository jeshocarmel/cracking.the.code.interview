package ch01

import "strings"

func isSubString(str1, str2 string) bool {

	newstr1 := []rune{}
	newstr1 = append(newstr1, []rune(str1)...)
	newstr1 = append(newstr1, []rune(str1)...)

	return strings.Contains(string(newstr1), str2)
}
