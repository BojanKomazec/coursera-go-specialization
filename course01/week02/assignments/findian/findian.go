// To build it use:
//   go build -o findian.exe findian.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// IsFound returns true if string starts with character "i", has somewhere "a" and ends with "n" while ignoring
// the trailing invisible CR and/or LF characters.
func IsFound(s string) bool {
	isFound := false
	str := strings.ToLower(s)

	// if exists, remove trailing Line Feed (added on Windows and on Unix/OSX)
	if strings.HasSuffix(str, "\n") {
		str = strings.TrimSuffix(str, "\n")
	}

	// if exist, remove trailing Carriage Return (added on Windows)
	if strings.HasSuffix(str, "\r") {
		str = strings.TrimSuffix(str, "\r")
	}

	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		isFound = true
	}

	return isFound
}

func main() {
	fmt.Println("Please enter a string: ")

	// All variants of fmt.ScanXXX place only characters before first SPACE into a string variable so we have to use
	// buffer.
	// When user presses ENTER, that adds End Of Line (CRLF (\r\n) on Windows and LF(\n) on Unix/OSX) at the end of the
	// string!
	stdinReader := bufio.NewReader(os.Stdin)

	if inputString, err := stdinReader.ReadString('\n'); err != nil {
		fmt.Println("Invalid input. Error: ", err)
	} else {
		if IsFound(inputString) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not found!")
		}
	}
}
