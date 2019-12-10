package ch15

import "time"
import "sync"
import "fmt"

import "math/rand"

const hunger = 3
const think = time.Second / 100
const eat = time.Second / 100

var dining sync.WaitGroup

// credits:  https://www.golangprograms.com/illustration-of-the-dining-philosophers-problem-in-golang.html

func diningProblem(philosopher string, dominantHand, otherHand *sync.Mutex) {
	fmt.Printf("%s is seated\n", philosopher)

	for h := hunger; h > 0; h-- {
		fmt.Printf("%s is hungry\n", philosopher)

		dominantHand.Lock()
		otherHand.Lock()
		fmt.Printf("%s is eating\n", philosopher)
		rand.Seed(time.Now().UnixNano())
		time.Sleep(eat + time.Duration(rand.Int63n(int64(eat))))

		// you can comment these two statements for a deadlock to occur
		dominantHand.Unlock()
		otherHand.Unlock()

		fmt.Printf("%s is thinking \n", philosopher)
		time.Sleep(think)
	}
	fmt.Printf("%s is satisified\n", philosopher)
	dining.Done()
	fmt.Printf("%s has left the table\n", philosopher)
}
