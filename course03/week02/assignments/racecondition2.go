package main

import (
	"fmt"
	"sync"
	"time"
)

var x int

func incrementX() {
	x = x + 1
}

func printX() {
	fmt.Println("x = ", x)
}

func waitForInput() {
	var dummyInput string
	fmt.Scanln(&dummyInput)
}

func attempt1() {
	go incrementX()
	go printX()
	waitForInput()
}

func attempt2() {
	for range [10]int{} {
		go incrementX()
		go printX()
	}

	waitForInput()
}

func attempt3() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		x = x + 1
		defer wg.Done()
	}()

	go func() {
		fmt.Println("x =", x)
		defer wg.Done()
	}()

	wg.Wait()
}

func attempt4() {
	go func() {
		for range [10]int{} {
			x = x + 1
		}
	}()

	go func() {
		fmt.Println("x =", x)
	}()

	waitForInput()
}

func attempt5() {
	var n int
	fmt.Scanln(&n)

	go func() {
		for i := 0; i < n; i++ {
			x = x + 1
		}
	}()

	go func() {
		fmt.Println("x =", x)
	}()

	waitForInput()
}

func attempt6() {
	go func() {
		x = x + 1
	}()
	go func() {
		fmt.Println("x =", x)
	}()
	time.Sleep(1 * time.Second)
}

// for /l %x in (1, 1, 100) do go run racecondition.go
func main() {
	attempt6()
}
