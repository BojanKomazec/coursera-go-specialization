package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type nameFormat struct {
	fname string
	lname string
}

func main() {

	var fileName string
	nameSlice := make([]nameFormat, 0, 20)
	fmt.Println("Enter the file name: ")

	fmt.Scan(&fileName)
	fholder, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer fholder.Close()
	sc := bufio.NewScanner(fholder)
	for sc.Scan() {
		line := sc.Text()
		words := strings.Fields(line)
		nameSlice = append(nameSlice, nameFormat{words[0], words[1]})
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	for _, name := range nameSlice {
		fmt.Println(name.fname, name.lname)
	}

}
