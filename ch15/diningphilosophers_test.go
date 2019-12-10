package ch15

import (
	"sync"
	"testing"
)

func TestDiningPhilosophers(t *testing.T) {

	t.Log("Table is empty")
	var philosophers = []string{"Socrates", "Sigmund Freud", "Daniel Kahneman", "David Hume", "Jesho Carmel"}
	dining.Add(5)

	fork0 := &sync.Mutex{}
	forkLeft := fork0

	for i := 1; i < len(philosophers); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}

	go diningProblem(philosophers[0], fork0, forkLeft)

	dining.Wait()
	t.Log("Table is empty")

}
