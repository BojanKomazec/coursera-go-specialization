package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalProperties struct {
	locomotion string
	food       string
	sound      string
}

type Callback func(string, string)
type AnimalMap map[string]Animal
type TypeMap map[string]AnimalProperties

type Cow struct {
	props *AnimalProperties
}

func (c Cow) Eat()   { fmt.Println(c.props.food) }
func (c Cow) Move()  { fmt.Println(c.props.locomotion) }
func (c Cow) Speak() { fmt.Println(c.props.sound) }

type Bird struct {
	props *AnimalProperties
}

func (b Bird) Eat()   { fmt.Println(b.props.food) }
func (b Bird) Move()  { fmt.Println(b.props.locomotion) }
func (b Bird) Speak() { fmt.Println(b.props.sound) }

type Snake struct {
	props *AnimalProperties
}

func (s Snake) Eat()   { fmt.Println(s.props.food) }
func (s Snake) Move()  { fmt.Println(s.props.locomotion) }
func (s Snake) Speak() { fmt.Println(s.props.sound) }

type Command interface {
	Execute(store AnimalMap, types TypeMap, params ...string)
}

type CommandNew struct {
	allowedParams []string
}
type CommandQuery struct {
	allowedParams []string
}

func (c *CommandNew) Execute(store AnimalMap, types TypeMap, params ...string) {
	checklParams(params, c.allowedParams, func(name, animalType string) {
		if animal := store[name]; animal == nil {
			props := types[animalType]
			switch animalType {
			case "cow":
				store[name] = Cow{&props}
			case "bird":
				store[name] = Bird{&props}
			case "snake":
				store[name] = Snake{&props}
			default:
				fmt.Println("Unknown animal", animalType)
			}

		} else {
			fmt.Printf("Animal with name %s already exists\n", name)
		}
	})
}

func (c *CommandQuery) Execute(store AnimalMap, types TypeMap, params ...string) {
	checklParams(params, c.allowedParams, func(name, info string) {
		if animal := store[name]; animal != nil {
			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown animal info", info)
			}
		} else {
			fmt.Println("Unknown animal with name", name)
		}
	})
}

func Contains(vs []string, t string) bool {
	exists := false
	for _, v := range vs {
		if v == t {
			exists = true
			break
		}
	}

	return exists
}

func checklParams(params []string, values []string, do Callback) {
	if len(params) >= 2 && Contains(values, params[1]) {
		do(params[0], params[1])
	} else {
		fmt.Println("Params are not valid")
	}
}

func main() {
	store := make(map[string]Animal)

	commandList := map[string]Command{
		"newanimal": &CommandNew{[]string{"cow", "bird", "snake"}},
		"query":     &CommandQuery{[]string{"eat", "move", "speak"}},
	}

	types := TypeMap{
		"cow":   AnimalProperties{"walk", "grass", "moo"},
		"bird":  AnimalProperties{"fly", "worms", "peep"},
		"snake": AnimalProperties{"slither", "mice", "hsss"},
	}

	for {
		var commandName, param1, param2 string

		fmt.Print(">")
		fmt.Scanf("%s %s %s", &commandName, &param1, &param2)

		if command := commandList[commandName]; command != nil {
			params := []string{param1, param2}
			command.Execute(store, types, params...)
		} else {
			fmt.Println("Unknown command")
		}

	}

}
