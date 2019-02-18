package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname [20]rune // string of size 20 (characters)
	lname [20]rune // string of size 20 (characters)
}

func trimTrailingCrLf(s string) string {
	trimmed := s
	if strings.HasSuffix(trimmed, "\n") {
		trimmed = strings.TrimSuffix(trimmed, "\n")
	}
	if strings.HasSuffix(trimmed, "\r") {
		trimmed = strings.TrimSuffix(trimmed, "\r")
	}

	return trimmed
}

// Using buffer reader allows both name and address to contain SPACE characters.
func inputFileName() string {
	var fileName string
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please enter the file name:")
		if name, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			fileName = trimTrailingCrLf(name)
			break
		}
	}

	return fileName
}

func main() {
	fileName := inputFileName()
	if file, err := os.Open(fileName); err != nil {
		fmt.Println("Error while opening a file:", err)
		os.Exit(1)
	} else {
		namesSplice := make([]name, 0)
		reader := bufio.NewReader(file)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			// fmt.Println(line)
			parts := strings.Split(line, " ")
			nameUntrimmed := parts[0]
			surnameUntrimmed := parts[1]
			// fmt.Println("Full name:", nameUntrimmed)
			// fmt.Println("Full surname:", surnameUntrimmed)

			runesName := []rune(nameUntrimmed)
			//var nameTrimmed string
			var nameFromLine name
			if len(runesName) > 20 {
				//nameTrimmed = string(runesName[0:20])
				// nameFromLine.fname = runesName[0:20]
				copy(nameFromLine.fname[:], runesName[0:20])
			} else {
				// nameTrimmed = string(runesName)
				// nameFromLine.fname = runesName
				copy(nameFromLine.fname[:], runesName)
			}

			runesSurname := []rune(surnameUntrimmed)
			//var surnameTrimmed string
			if len(runesSurname) > 20 {
				// surnameTrimmed = string(runesSurname[0:20])
				// nameFromLine.lname = runesSurname[0:20]
				copy(nameFromLine.lname[:], runesSurname[0:20])
			} else {
				// surnameTrimmed = string(runesSurname)
				// nameFromLine.lname = runesSurname
				copy(nameFromLine.lname[:], runesSurname)
			}

			// fmt.Println("Trimmed name:", string(nameFromLine.fname[:]))
			// fmt.Println("Trimmed surname:", string(nameFromLine.lname[:]))

			namesSplice = append(namesSplice, nameFromLine)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		fmt.Println("\nNames (splice elements trimmed names to max 20 characters):")
		fmt.Println()
		for index := 0; index < len(namesSplice); index++ {
			fmt.Println("Name:", string(namesSplice[index].fname[:]))
			fmt.Println("Surname:", string(namesSplice[index].lname[:]))
			fmt.Println()
		}
	}
}
