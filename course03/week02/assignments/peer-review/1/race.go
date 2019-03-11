package main

import (
	"fmt"
	// "math/rand"
	"time"
)

/* A race condition exists between the following 2 goroutines due to the shared variable `i`. It can occur when i is printed from the second goroutine when the first one is still calculating the final value. */

func main() {
	var i int

	go func() {
		i = 5
		//time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		i = i + 4
	}()

	go func() {
		//time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
		fmt.Println(i)
	}()

	/* sleep for a bit so that the results are actually printed out */
	time.Sleep(100 * time.Millisecond)

}
