package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var animalMap = make(map[string]Animal, 0)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type cow struct {
}

func (c cow) Speak() {
	fmt.Println("moo")
}

func (c cow) Eat() {
	fmt.Println("grass")
}
func (c cow) Move() {
	fmt.Println("walk")
}

type bird struct {
}

func (b bird) Eat() {
	fmt.Println("worms")
}

func (b bird) Move() {
	fmt.Println("fly")
}

func (b bird) Speak() {
	fmt.Println("peep")
}

type snake struct {
}

func (s snake) Eat() {
	fmt.Println("mice")
}

func (s snake) Move() {
	fmt.Println("slither")
}

func (s snake) Speak() {
	fmt.Println("hsss")
}

func MakeAnimal(name string) error {
	switch name {
	case "cow":
		if _, ok := animalMap[name]; !ok {
			animalMap[name] = cow{}
			return nil
		} else {
			return errors.New("cow is already made")
		}
	case "bird":
		if _, ok := animalMap[name]; !ok {
			animalMap[name] = bird{}
			return nil
		} else {
			return errors.New("bird is already made")
		}
	case "snake":
		if _, ok := animalMap[name]; !ok {
			animalMap[name] = snake{}
			return nil
		} else {
			return errors.New("snake is already made")
		}
	default:
		return errors.New("Can not make such animal")

	}
}
func QueryAnimal(name string, cmd string) error {
	if v, ok := animalMap[name]; ok {
		switch cmd {
		case "eat":
			v.Eat()
			return nil
		case "move":
			v.Move()
			return nil
		case "speak":
			v.Speak()
			return nil
		default:
			return errors.New("invalid command")
		}
	} else {
		return errors.New("No such animal")
	}
}
func main() {
	for {
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		line = strings.ToLower(line)
		cmd := strings.Split(line, " ")
		switch cmd[0] {
		case "newanimal":
			if err := MakeAnimal(cmd[1]); err != nil {
				fmt.Println(err)
				continue
			}
		case "query":
			if err := QueryAnimal(cmd[1], cmd[2]); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

}
