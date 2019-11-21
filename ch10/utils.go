package ch10

import (
	"math/rand"
	"time"
)

func generateArray(n int) []int {

	arr := make([]int, n)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}

	return arr
}
