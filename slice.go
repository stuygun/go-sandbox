package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 3)
	message := "Please enter a number or enter 'X' to exit.."
	validInput := 0
	for {
		fmt.Println(message)
		var input string
		_, err := fmt.Scan(&input)
		if err == nil {
			if input != "X" {
				num, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Please enter a valid number or enter 'X' to exit!")
					continue
				} else {
					validInput++
					if validInput > 3 {
						slice = append(slice, num)
					} else {
						slice[validInput-1] = num
					}
				}
			} else {
				sort.Ints(slice)
				fmt.Println("Slice: ", slice)
				break
			}
		}
	}
}
