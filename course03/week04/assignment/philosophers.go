// Bojan Komazec, 2019
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	id             int
	leftChopstick  *chopstick
	rightChopstick *chopstick
}

func (p philosopher) eat(wg *sync.WaitGroup, currentEatersCountChannel chan struct{}) {
	// fmt.Println("Philosopher", p.id, "eat()")
	for i := 0; i < 3; i++ {
		currentEatersCountChannel <- struct{}{}

		// randomize picking chopsticks
		randBinaryValue := rand.Intn(2)
		// fmt.Println(randBinaryValue)
		if randBinaryValue == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		fmt.Println("			Philosopher", p.id, ": Starting to eat - round", i+1)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("			Philosopher", p.id, ": Finishing eating - round", i+1)

		if randBinaryValue == 0 {
			p.rightChopstick.Unlock()
			p.leftChopstick.Unlock()
		} else {
			p.leftChopstick.Unlock()
			p.rightChopstick.Unlock()
		}

		<-currentEatersCountChannel
	}
	// fmt.Println("Philosopher", p.id, "~eat()")
	wg.Done()
}

// One chopstick is placed between two adjacent philosophers.
// Philosopher needs left and right chopstic in order to eat.
func main() {
	const PhilosophersCount int = 5
	const MaxConcurrentEaters int = 2

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// A "host" mentioned in the assignment text is actually main() function - which IS executed in its own goroutine :)
	// It gives permissions to philosophers to eat via this channel which has a buffer of size 2 so only 2 philosophers
	// can actually eat (start locking chopsticks...) at the same time.
	currentEatersCountChannel := make(chan struct{}, MaxConcurrentEaters)

	chopsticks := make([]*chopstick, PhilosophersCount)
	for i := 0; i < PhilosophersCount; i++ {
		chopsticks[i] = new(chopstick)
	}

	wg.Add(PhilosophersCount)

	philosophers := make([]*philosopher, PhilosophersCount)
	for i := 0; i < PhilosophersCount; i++ {
		philosophers[i] = &philosopher{
			i + 1,
			chopsticks[i],
			chopsticks[(i+1)%PhilosophersCount],
		}
		go philosophers[i].eat(&wg, currentEatersCountChannel)
	}

	wg.Wait()
}
