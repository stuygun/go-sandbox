package main

import (
	"fmt"
	"strings"
)

type Animal1 struct {
	animalName       string
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

func (a Animal1) Name() string {
	return a.animalName
}

func (a Animal1) Eat() {
	fmt.Println(a.foodEaten)
}

func (a Animal1) Move() {
	fmt.Println(a.locomotionMethod)
}

func (a Animal1) Speak() {
	fmt.Print(a.spokenSound)
}

func main() {
	cow := Animal1{"cow", "grass", "walk", "moo"}
	bird := Animal1{"bird", "worms", "fly", "peep"}
	snake := Animal1{"snake", "mice", "slither", "hsss"}

	var animalMap = make(map[string]Animal1)
	animalMap["cow"] = cow
	animalMap["bird"] = bird
	animalMap["snake"] = snake

	message := "Please enter name of animal (cow, bird, or snake) and action (eat, move, or speak) in a Comma-Separated-Form (e.g. cow,eat)\n>"
	for {
		fmt.Print(message)
		var input string
		_, err := fmt.Scan(&input)
		if err == nil {
			split := strings.Split(input, ",")
			if len(split) != 2 {
				continue
			} else {
				name := strings.ToLower(strings.TrimSpace(split[0]))
				action := strings.ToLower(strings.TrimSpace(split[1]))

				animal, exists := animalMap[name]
				if !exists {
					fmt.Println("Animal1 (" + name + ") not found! Available animals are cow, bird or snake!")
					continue
				} else {
					switch action {
					case "eat":
						animal.Eat()
					case "move":
						animal.Move()
					case "speak":
						animal.Move()
					default:
						fmt.Println("Action (" + action + ") not found! Available actions are eat,move or speak!")
						continue
					}
				}

			}
		}
	}
}
