package ch15

import (
	"fmt"
	"sync"
)

var mux1 sync.Mutex
var mux2 sync.Mutex

var cioWg sync.WaitGroup

// as the problem only requires a sequential execution of methods i.e. one at a time \
// it's safe to use a mutex instead of a semaphore.

//CallInOrder calls the three methods first() second() third() using goroutines.
func CallInOrder() string {
	cioWg.Add(3)

	mux1.Lock()
	mux2.Lock()

	go first()
	go second()
	go third()

	cioWg.Wait()
	return "method executions completed"
}

func first() {
	fmt.Println("This is first method")
	mux1.Unlock()
	cioWg.Done()
}

func second() {
	mux1.Lock()
	mux1.Unlock()
	fmt.Println("This is second method")
	mux2.Unlock()

	cioWg.Done()
}

func third() {
	mux2.Lock()
	fmt.Println("This is third method")
	mux2.Unlock()
	cioWg.Done()
}
