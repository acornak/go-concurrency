package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// constants
const hunger = 3

// variables
var philosophers = []string{"Plato", "Socrates", "Aristoteles", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 3 * time.Second
var thinkTime = 2 * time.Second
var orderFinished []string
var orderMutex sync.Mutex

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	// print a message
	fmt.Printf("%s is seated.\n", philosopher)
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Printf("%s is hungry.\n", philosopher)
		time.Sleep(sleepTime)

		// lock both forks
		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left.\n", philosopher)

		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right.\n", philosopher)

		// print a message
		fmt.Printf("%s has both forks, and is eating.\n", philosopher)
		time.Sleep(eatTime)

		// give the philosopher some time to think
		fmt.Println(philosopher, "is thinking.")
		time.Sleep(thinkTime)

		// unlock the mutexes
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left.\n", philosopher)

		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right.\n", philosopher)
		time.Sleep(sleepTime)
	}

	// print out done message
	fmt.Printf("%s is satisfied.\n", philosopher)
	time.Sleep(sleepTime)

	fmt.Printf("%s has left the table.\n", philosopher)

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()
}

func main() {
	// print intro
	fmt.Println("-------------------------------")
	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("-------------------------------")

	wg.Add(len(philosophers))

	// create a mutex for the left fork
	forkLeft := &sync.Mutex{}

	// spawn 1 go routine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the right fork
		forkRight := &sync.Mutex{}

		// call go routine
		go diningProblem(philosophers[i], forkLeft, forkRight)

		// not a copy!
		forkLeft = forkRight
	}

	wg.Wait()

	fmt.Println("The table is empty.")
	fmt.Println("-------------------")
	fmt.Printf("Order finished: %s\n", strings.Join(orderFinished, ", "))
	fmt.Println("-------------------")
}
