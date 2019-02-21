package main

import (
	"errors"
	"fmt"
)

// Named interface describes any entity that can have a name
type Named interface {
	getName() string
}

// Animal interface contains methods that print what animal eat, how it moves and how it speaks.
type Animal interface {
	Named
	Eat()
	Move()
	Speak()
}

// Cow type
type Cow struct {
	name string
}

// Bird type
type Bird struct {
	name string
}

// Snake type
type Snake struct {
	name string
}

//
// Named interface
//

func (cow Cow) getName() string {
	return cow.name
}

func (bird Bird) getName() string {
	return bird.name
}

func (snake Snake) getName() string {
	return snake.name
}

//
// Animal interface
//

// Eat method from Animal interface
func (cow Cow) Eat() {
	fmt.Println("grass")
}

// Eat method from Animal interface
func (bird Bird) Eat() {
	fmt.Println("worms")
}

// Eat method from Animal interface
func (snake Snake) Eat() {
	fmt.Println("mice")
}

// Move method from Animal interface
func (cow Cow) Move() {
	fmt.Println("walk")
}

// Move method from Animal interface
func (bird Bird) Move() {
	fmt.Println("fly")
}

// Move method from Animal interface
func (snake Snake) Move() {
	fmt.Println("slither")
}

// Speak method from Animal interface
func (cow Cow) Speak() {
	fmt.Println("moo")
}

// Speak method from Animal interface
func (bird Bird) Speak() {
	fmt.Println("peep")
}

// Speak method from Animal interface
func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func isValidCommand(command string) bool {
	return command == "newanimal" || command == "query"
}

func isValidAnimal(animal string) bool {
	return animal == "cow" || animal == "bird" || animal == "snake"
}

func isValidAnimalProperty(property string) bool {
	return property == "eat" || property == "move" || property == "speak"
}

func getRequest() (string, string, string) {
	var command, name, option string
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&command, &name, &option); err != nil {
			fmt.Println("Error:", err)
		} else {
			if isValidCommand(command) {
				if command == "newanimal" {
					if isValidAnimal(option) {
						break
					} else {
						fmt.Println("Invalid animal type. Please use 'cow', 'bird' or 'snake'.")
					}
				} else if command == "query" {
					if isValidAnimalProperty(option) {
						break
					} else {
						fmt.Println("Invalid animal property. Please use 'eat', 'move' or 'speak'.")
					}
				}
			} else {
				fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
			}
		}
	}
	return command, name, option
}

func appendAnimal(animals *[]Animal, animal Animal) {
	*animals = append(*animals, animal)
	fmt.Println("Created it!")
}

func findAnimal(animals []Animal, name string) (Animal, error) {
	var foundAnimal Animal
	var err = errors.New("Animal with that name not found")
	for _, animal := range animals {
		if Named(animal).getName() == name {
			foundAnimal = animal
			err = nil
			break
		}
	}
	return foundAnimal, err
}

func main() {
	var animals []Animal

	for {
		command, name, option := getRequest()
		switch command {
		case "newanimal":
			switch option {
			case "cow":
				appendAnimal(&animals, Cow{name})
			case "bird":
				appendAnimal(&animals, Bird{name})
			case "snake":
				appendAnimal(&animals, Snake{name})
			}
		case "query":
			var animal Animal
			var err error
			if animal, err = findAnimal(animals, name); err != nil {
				fmt.Println(err)
				break
			}
			switch option {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	}
}
