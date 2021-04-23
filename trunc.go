package main

import (
	"fmt"
	"strconv"
)

func main() {
	message := "Please enter a number or enter 'q' to exit.."
	for {
		fmt.Println(message)
		var input string
		_, err := fmt.Scan(&input)
		if err == nil {
			if input != "q" {
				num, err := strconv.ParseFloat(input, 64)
				trunc := int64(num)
				fmt.Printf("truncated: %d\n", trunc)
				if err != nil {
					fmt.Println("Please enter a valid number or enter 'q' to exit!")
					continue
				}
			} else {
				fmt.Println("Good Bye!")
				break
			}
		}
	}
}
