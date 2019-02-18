package main

import "fmt"

type Animal struct {
	locomotion string
	food       string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := make(map[string]Animal)
	animals["cow"] = Animal{locomotion: "walk", food: "grass", noise: "moo"}
	animals["bird"] = Animal{locomotion: "fly", food: "worms", noise: "peep"}
	animals["snake"] = Animal{locomotion: "slither", food: "mice", noise: "hsss"}

	for {
		var name, info string
		fmt.Print("> ")
		fmt.Scanf("%s %s", &name, &info)

		animal, ok := animals[name]

		if !ok {
			fmt.Println("Unknown animal", name)
			continue
		}

		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("No info named '%s' about animal %s\n", info, name)
		}
	}

}
