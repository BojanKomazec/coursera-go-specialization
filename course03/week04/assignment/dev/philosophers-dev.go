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

func (p philosopher) eat(wg *sync.WaitGroup, currentEatersCountChannel chan int) {
	fmt.Println("Philosopher", p.id, "eat()")
	for i := 0; i < 3; i++ {
		currentEatersCountChannel <- p.id

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
	fmt.Println("Philosopher", p.id, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~eat()")
	wg.Done()
}

// func host(wg *sync.WaitGroup) {

// 	wg.Done()
// }

// one chopstic is between two adjacent philosophers; philosopher needs left and right chopstic
// in order to eat
func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	currentEatersCountChannel := make(chan int, 2)

	chopsticks := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(chopstick)
	}

	wg.Add(5)

	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{
			i + 1,
			chopsticks[i],
			chopsticks[(i+1)%5],
		}
		go philosophers[i].eat(&wg, currentEatersCountChannel)
	}

	// for i := 0; i < 5; i++ {
	// }

	// go host(&wg)
	wg.Wait()
}

// type Philosopher struct {
// 	id                int
// 	inboundMsgChannel chan struct{}
// }

// func (p Philosopher) eat(wg *sync.WaitGroup, requestsChannel chan int, currentEatersCountChannel chan int) {
// 	fmt.Println("Philosopher", p.id, "eat()")
// 	for i := 0; i < 3; {

// 		select {
// 		// Wait for approval by the host (host will send struct{}{} once approved).
// 		// Blocking if channel is empty
// 		case <-p.inboundMsgChannel:
// 			// Once approved, notify host that one more philosopher started eating
// 			//currentEatersCountChannel <- p.id
// 			fmt.Println("			Philosopher", p.id, ": Starting to eat")
// 			time.Sleep(100 * time.Millisecond)
// 			fmt.Println("			Philosopher", p.id, ": Finishing eating")
// 			i = i + 1
// 			// notify host that ther is one less philosopher eating
// 			<-currentEatersCountChannel
// 		case requestsChannel <- p.id:
// 			fmt.Println("Philosopher", p.id, ": sent request")
// 			// Send request for eating to host.
// 			// Blocks if channel is full (there are alredy 2 philosophers requests in the queue).
// 			// fmt.Println("			Philosopher", p.id, ": NOP")
// 		}

// 	}
// 	fmt.Println("Philosopher", p.id, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~eat()")
// 	wg.Done()
// }

// func absInt(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

// host can't accept random requests as it can happend that philosophers 1, 3 and then 2 finish eating completely
// leaving 4 and 5 to eat but they can't eat at the same time
// func host(wg *sync.WaitGroup, requestsChannel chan int, currentEatersCountChannel chan int, philosophers []*Philosopher) {
// 	var previousAcceptedEaterID int
// 	var totalEatings int

// 	//for totalEatings < 15 { // loop till you receive the request from some philosopher
// 	for {
// 		select {
// 		case id := <-requestsChannel: // blocking if channel is empty
// 			fmt.Println("host(): received request from", id)

// 			var diff int
// 			if previousAcceptedEaterID != 0 {
// 				diff = absInt(id - previousAcceptedEaterID)
// 			}

// 			if previousAcceptedEaterID == 0 || (previousAcceptedEaterID != 0 && ((diff == 2) || (diff == 3) || (diff == 0))) {
// 				fmt.Println("host(): accepted request from", id)
// 				previousAcceptedEaterID = id

// 				// this is blocking if buffer is full (if 2 philosophers are already eating)
// 				currentEatersCountChannel <- id
// 				fmt.Println("host(): about to unblock", id)
// 				philosophers[id-1].inboundMsgChannel <- struct{}{} // blocking if channel is full
// 				fmt.Println("host(): unblocked", id)
// 				totalEatings = totalEatings + 1
// 				fmt.Println("host(): totalEatings =", totalEatings)
// 			} else {
// 				fmt.Println("host(): !!!!!!!!!!!!!!!!!!!!! REFUSED request from", id)
// 			}
// 		default:
// 			fmt.Println("host(): no requests")
// 			if totalEatings >= 15 {
// 				wg.Done()
// 				return
// 			}
// 		}
// 	}

// }

// func main() {
// 	var wg sync.WaitGroup

// 	requestsChannel := make(chan int, 1)
// 	currentEatersCountChannel := make(chan int, 2)

// 	philosophers := make([]*Philosopher, 5)
// 	for i := 0; i < 5; i++ {
// 		philosophers[i] = &Philosopher{
// 			(i + 1),
// 			make(chan struct{}),
// 		}
// 	}
// 	wg.Add(16)
// 	for i := 0; i < 5; i++ {
// 		go philosophers[i].eat(&wg, requestsChannel, currentEatersCountChannel)
// 	}

// 	go host(&wg, requestsChannel, currentEatersCountChannel, philosophers)

// 	wg.Wait()
// }

// host executes in its own goroutine
// In order to allow both philosophers to eat, they should not be sitting next to each other.
// E.g. if Philospher 1 can eat at the same time as philosopher 3 or 4 but not 2 or 5.
// The simplest pattern host can enforce is Round Robin:
// 1-3, 2-4, 3-5, 4-1, 5-2, 1-3, 2-4, 5
// It can start filling the semaphore channel with those values, two by two and
// go func() {
// 	fmt.Println("host()")

// 	mtx.Lock()
// 	channel <- 1
// 	channel <- 3
// 	mtx.Unlock()
// 	//wg.Wait()
// 	mtx.Lock()
// 	channel <- 2
// 	channel <- 4
// 	mtx.Unlock()
// 	wg.Wait()
// 	mtx.Lock()
// 	channel <- 3
// 	channel <- 5
// 	mtx.Unlock()
// 	wg.Wait()
// 	mtx.Lock()
// 	channel <- 4
// 	channel <- 1
// 	mtx.Unlock()
// 	wg.Wait()
// 	mtx.Lock()
// 	channel <- 5
// 	channel <- 2
// 	mtx.Unlock()
// 	wg.Wait()
// 	mtx.Lock()
// 	channel <- 1
// 	channel <- 3
// 	mtx.Unlock()
// 	wg.Wait()
// 	mtx.Lock()
// 	channel <- 2
// 	channel <- 4
// 	mtx.Unlock()
// 	wg.Wait()
// 	mtx.Lock()
// 	channel <- 5
// 	mtx.Unlock()
// 	wg.Wait()

// 	fmt.Println("~host()")
// 	wgMain.Done()
// }()
