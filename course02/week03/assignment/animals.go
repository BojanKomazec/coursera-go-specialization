package main

import (
	"fmt"
)

// Animal struct defines an animal
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints animal's food
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

// Move prints animal's locomotion
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

// Speak prints animal's noise
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func getRequest() (string, string) {
	var animalChoice, infoChoice string
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&animalChoice, &infoChoice); err != nil {
			fmt.Println("Error:", err)
		} else {
			break
		}
	}
	return animalChoice, infoChoice
}

func main() {
	fmt.Println("Please type '[cow|bird|snake] [eat|move|speak]' after the prompt \">\" to discover animals. Press CTRL+C to exit.")
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		animalChoice, infoChoice := getRequest()

		var animal Animal
		if animalChoice == "cow" {
			animal = cow
		} else if animalChoice == "bird" {
			animal = bird
		} else if animalChoice == "snake" {
			animal = snake
		}

		if infoChoice == "eat" {
			animal.Eat()
		} else if infoChoice == "move" {
			animal.Move()
		} else if infoChoice == "speak" {
			animal.Speak()
		}
	}
}
