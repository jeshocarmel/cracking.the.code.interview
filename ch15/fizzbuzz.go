package ch15

import (
	"strconv"
	"sync"
)

//FizzBuzzSingleThreaded ...
func FizzBuzzSingleThreaded(n int) map[int]string {

	var result = make(map[int]string)

	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			result[i] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i] = "Fizz"
		} else if i%5 == 0 {
			result[i] = "Buzz"
		} else {
			result[i] = strconv.Itoa(i)
		}
	}

	return result
}

//FizzBuzzRoutines ...
func FizzBuzzRoutines(n int) map[int]string {

	var result = make(map[int]string)

	var wg sync.WaitGroup
	wg.Add(n * 4)

	go func() {
		for i := 1; i <= n; i++ {
			if i%3 == 0 && i%5 == 0 {
				result[i] = "FizzBuzz"
			}
			wg.Done()
		}
	}()

	go func() {
		for i := 1; i <= n; i++ {
			if i%3 == 0 && i%5 != 0 {
				result[i] = "Fizz"
			}
			wg.Done()
		}
	}()

	go func() {
		for i := 1; i <= n; i++ {
			if i%3 != 0 && i%5 == 0 {
				result[i] = "Buzz"
			}
			wg.Done()
		}
	}()

	go func() {
		for i := 1; i <= n; i++ {
			if i%3 != 0 && i%5 != 0 {
				result[i] = strconv.Itoa(i)
			}
			wg.Done()
		}
	}()

	wg.Wait()
	return result
}
