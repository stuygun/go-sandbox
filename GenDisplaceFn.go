package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	message := "Enter acceleration,initial velocity and initial displacement in Comma-Separated-Format(without any space e.g. 10,2,1)"
	var input string
	var acceleration, initialVelocity, initialDisplacement, time float64
	var err error

	for {
		fmt.Println(message)
		_, _ = fmt.Scan(&input)
		split := strings.Split(input, ",")
		if len(split) != 3 {
			fmt.Println(message)
			continue
		} else {
			acceleration, err = strconv.ParseFloat(split[0], 64)
			if err != nil {
				fmt.Println(message)
				continue
			}
			initialVelocity, err = strconv.ParseFloat(split[1], 64)
			if err != nil {
				fmt.Println(message)
				continue
			}
			initialDisplacement, err = strconv.ParseFloat(split[2], 64)
			if err != nil {
				fmt.Println(message)
				continue
			}
			break
		}
	}
	for {
		fmt.Println("Enter the time:")
		_, _ = fmt.Scan(&input)
		time, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(message)
			continue
		}
		break
	}

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Print("Displacement: ")
	fmt.Print(fn(time))
	fmt.Println()
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	function := func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
	}
	return function
}
