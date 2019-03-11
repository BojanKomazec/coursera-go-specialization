package main

import (
	"fmt"
	"sync"
	"time"
)

type Chops struct {
	sync.Mutex
}

type Philos struct {
	left   *Chops
	right  *Chops
	number int
}

func (p *Philos) Eat(wg *sync.WaitGroup, host chan int) {
	for i := 0; i < 3; i++ {
		host <- p.number
		p.left.Lock()
		p.right.Lock()

		fmt.Println("starting to eat <", p.number, ">")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("finishing eating <", p.number, ">")

		p.right.Unlock()
		p.left.Unlock()
		<-host
	}
	wg.Done()
}

func main() {
	chops := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(Chops)
	}

	philos := make([]*Philos, 5)
	for i := 0; i < 5; i++ {
		philos[i] = new(Philos)
		philos[i].number = i + 1
		philos[i].left = chops[i]
		philos[i].right = chops[(i+1)%5]
	}

	var wg sync.WaitGroup
	wg.Add(5)

	host := make(chan int, 2)
	for i := 0; i < 5; i++ {
		go philos[i].Eat(&wg, host)
	}
	wg.Wait()
}
