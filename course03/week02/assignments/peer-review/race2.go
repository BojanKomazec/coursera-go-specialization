/*
A race condition occurs when the result of the computation is dependent on the execution order
of instructions. In other words, the result depends on interleavings of instructions
(as Ian Harris calls it).

In a concurrent program the execution order of instructions from multiple goroutines is decided by
the Go Runtime Scheduler. This order is not deterministic, and may differ between runs of the program.
If two (or more) goroutines are communicating using same resource, then the resource one goroutine is
using for computation can be modified by other goroutines. This can lead to different results in each
program execution.

In this program there are two goroutines which have the same global variable in scope. The global
variable is initialized to 10. One goroutine just checks whether the variable is equal to 10, and
prints appropriate message. The other goroutine only increments the global variable.

When the two goroutines are running concurrently, it sometimes happens that the global variable
is incremented by one goroutine before it is checked by the other, even though the goroutine that
checks the value is started first in main().

You may have to run this program many times to actually observe that (in my testing the probability
of such race was about 3%). Inserting some random delays into both goroutines would increase the
probability, but I wanted to keep the code simple.

In Linux when using Bash shell:
   while true; do go run race.go; done
*/

package main

import (
	"fmt"
	"time"
)

var global int = 10

func checkGlobal() {
	if global == 10 {
		fmt.Println("global is 10")
	} else {
		fmt.Println("global is not 10!")
	}
}

func incrementGlobal() {
	global++
}

func main() {
	go checkGlobal()
	go incrementGlobal()

	/* sleep for a bit so that the results are actually printed out */
	time.Sleep(200 * time.Millisecond)
}
