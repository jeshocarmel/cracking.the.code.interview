package ch01

import (
	"strconv"
	"strings"
)

func compressString(str string) string {

	count := 1
	var sb strings.Builder
	for i := 0; i+1 < len(str); i++ {

		if str[i+1] == str[i] {
			count++
		} else {
			sb.WriteByte(str[i])
			sb.WriteString(strconv.Itoa(count))
			count = 1
		}
	}

	sb.WriteByte(str[len(str)-1])
	sb.WriteString(strconv.Itoa(count))

	if sb.Len() >= len(str) {
		return str
	}

	return sb.String()
}
